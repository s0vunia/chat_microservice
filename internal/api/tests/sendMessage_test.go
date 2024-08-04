package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/s0vunia/chat_microservice/internal/api/chat"
	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/chat_microservice/internal/service"
	"github.com/s0vunia/chat_microservice/internal/service/mocks"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
	"github.com/stretchr/testify/require"
)

func TestImplementation_SendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id         = gofakeit.UUID()
		fromUserID = gofakeit.Int64()
		toChatID   = gofakeit.Int64()
		text       = gofakeit.BeerName()

		serviceErr = fmt.Errorf("service error")

		req = &desc.SendMessageRequest{
			Message: &desc.MessageCreate{
				FromUserId: fromUserID,
				ToChatId:   toChatID,
				Text:       text,
			},
		}

		createMessage = model.MessageCreate{
			Info: model.MessageInfo{
				ChatID: toChatID,
				UserID: fromUserID,
				Text:   text,
			},
		}

		res = &desc.SendMessageResponse{
			Id:     id,
			ChatId: toChatID,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.SendMessageResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, &createMessage).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, &createMessage).Return("", serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sendMessageServiceMock := tt.chatServiceMock(mc)
			api := chat.NewImplementation(sendMessageServiceMock)

			newID, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
