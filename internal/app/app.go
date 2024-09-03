package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"sync"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/s0vunia/chat_microservice/internal/closer"
	"github.com/s0vunia/chat_microservice/internal/config"
	"github.com/s0vunia/chat_microservice/internal/interceptor"
	"github.com/s0vunia/chat_microservice/internal/logger"
	"github.com/s0vunia/chat_microservice/internal/metric"
	"github.com/s0vunia/chat_microservice/internal/tracing"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	configPath  string
	serviceName = "chat-service"
)

func init() {
	configPath = os.Getenv("CONFIG_PATH")
}

// App represents the app.
type App struct {
	serviceProvider  *serviceProvider
	grpcServer       *grpc.Server
	prometheusServer *http.Server
}

// NewApp creates a new app.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run runs the app.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			logger.Fatal(
				"failed to run GRPC server",
				zap.Error(err),
			)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runPrometheus()
		if err != nil {
			logger.Fatal(
				"failed to run prometheus server",
				zap.Error(err),
			)
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initLogger,
		a.initMetric,
		a.initTracing,
		a.initGRPCServer,
		a.initPrometheusServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.Init(a.getCore(a.getAtomicLevel()))
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initPrometheusServer(_ context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:              a.serviceProvider.PrometheusConfig().Address(),
		Handler:           mux,
		ReadHeaderTimeout: a.serviceProvider.PrometheusConfig().ReadTimeout(),
	}
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.MetricsInterceptor,
				interceptor.ServerTracingInterceptor,
				interceptor.LogInterceptor,
				interceptor.AuthInterceptor(a.serviceProvider.AuthService(ctx)),
			),
		),
	)

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImpl(ctx))

	return nil
}

func (a *App) initMetric(ctx context.Context) error {
	err := metric.Init(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initTracing(_ context.Context) error {
	tracing.Init(logger.Logger(), serviceName)
	return nil
}

func (a *App) runGRPCServer() error {
	logger.Info("GRPC server is running",
		zap.String("address", a.serviceProvider.GRPCConfig().Address()),
	)

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runPrometheus() error {
	logger.Info("Prometheus server is running",
		zap.String("address", a.serviceProvider.PrometheusConfig().Address()),
	)

	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   a.serviceProvider.LoggerConfig().FileName(),
		MaxSize:    a.serviceProvider.LoggerConfig().MaxSize(), // megabytes
		MaxBackups: a.serviceProvider.LoggerConfig().MaxBackups(),
		MaxAge:     a.serviceProvider.LoggerConfig().MaxAge(), // days
	})
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
}

func (a *App) getAtomicLevel() zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(a.serviceProvider.LoggerConfig().Level()); err != nil {
		logger.Fatal("failed to set log level", zap.Error(err))
	}

	return zap.NewAtomicLevelAt(level)
}
