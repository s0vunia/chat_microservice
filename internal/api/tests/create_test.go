package tests

import (
	"context"
	"fmt"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/api/chat"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/model"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/service"
	"github.com/s0vunia/chat_microservices_course_boilerplate/internal/service/mocks"
	desc "github.com/s0vunia/chat_microservices_course_boilerplate/pkg/chat_v1"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestImplementation_Create(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id   = gofakeit.Int64()
		name = gofakeit.Animal()

		ids          = []int64{1, 2, 3, 4, 5}
		participants = []model.ParticipantCreate{
			{UserID: 1},
			{UserID: 2},
			{UserID: 3},
			{UserID: 4},
			{UserID: 5},
		}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Chat: &desc.ChatCreate{
				Name: name,
			},
			UserIds: ids,
		}

		chatCreate = model.ChatCreate{
			Name: name,
		}
		participantsCreate = model.ParticipantsCreate{
			Participants: participants,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)
	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, &chatCreate, &participantsCreate).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, &chatCreate, &participantsCreate).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			createServiceMock := tt.chatServiceMock(mc)
			api := chat.NewImplementation(createServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
