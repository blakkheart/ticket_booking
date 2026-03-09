-- +goose Up
-- +goose StatementBegin
CREATE TABLE account (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    role VARCHAR(128) NOT NULL
);
CREATE TABLE event (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    date timestamptz,
    location VARCHAR(255) NOT NULL
);
CREATE TABLE ticket (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price bigint NOT NULL,
    quantity integer NOT NULL,
    event_id bigint references event(id) NOT NULL
);
CREATE TABLE booking (
    id BIGSERIAL PRIMARY KEY,
    account_id bigint references account(id) NOT NULL,
    event_id bigint references event(id) NOT NULL,
    quantity integer NOT NULL
);
CREATE TABLE refresh_token (
    user_id BIGSERIAL PRIMARY KEY ,
    token_hash VARCHAR(512) NOT NULL UNIQUE,
    expires_at timestamptz NOT NULL,
    revoked BOOLEAN NOT NULL,
    replaced_by BIGINT,
    FOREIGN KEY(user_id) REFERENCES account(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account CASCADE;
DROP TABLE event CASCADE;
DROP TABLE ticket CASCADE;
DROP TABLE booking CASCADE;
DROP TABLE refresh_token CASCADE;
-- +goose StatementEnd
