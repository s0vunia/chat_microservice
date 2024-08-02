-- +goose Up
-- +goose StatementBegin
CREATE TABLE ChatParticipants
(
    chat_id    INT REFERENCES Chats (id) ON DELETE CASCADE,
    user_id    INT       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (chat_id, user_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ChatParticipants;
-- +goose StatementEnd
