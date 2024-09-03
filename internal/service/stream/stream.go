package stream

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/model"
)

// Stream represents the stream.
type Stream interface {
	Send(*model.MessageCreate) error
	Context() context.Context
}
