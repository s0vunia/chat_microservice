package desc

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/converter"
	"github.com/s0vunia/chat_microservice/internal/model"
	"github.com/s0vunia/chat_microservice/internal/service/stream"
	desc "github.com/s0vunia/chat_microservice/pkg/chat_v1"
)

var _ stream.Stream = (*StreamDesc)(nil)

// StreamDesc implements stream.Stream
type StreamDesc struct {
	stream desc.ChatV1_ConnectChatServer
}

// NewStream returns new stream
func NewStream(stream desc.ChatV1_ConnectChatServer) *StreamDesc {
	return &StreamDesc{
		stream: stream,
	}
}

// Send sends message
func (s *StreamDesc) Send(msg *model.MessageCreate) error {
	return s.stream.Send(converter.ToMessageCreateFromModel(msg))
}

// Context returns context
func (s *StreamDesc) Context() context.Context {
	return s.stream.Context()
}
