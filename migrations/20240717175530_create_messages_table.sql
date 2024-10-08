-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE Messages
(
    id         UUID               DEFAULT uuid_generate_v4(),
    chat_id    INT REFERENCES Chats (id) ON DELETE CASCADE,
    user_id    INT       NOT NULL,
    text       TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    PRIMARY KEY (id, chat_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Messages;
-- +goose StatementEnd
