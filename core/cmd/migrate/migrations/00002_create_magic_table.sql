-- +goose Up
-- +goose StatementBegin
CREATE TABLE magic (
    email TEXT NOT NULL UNIQUE,
    code TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (email) REFERENCES users(email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE magic;
-- +goose StatementEnd
