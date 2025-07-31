-- +goose Up
-- +goose StatementBegin
CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    date timestamp,
    location VARCHAR(255) NOT NULL
);
CREATE TABLE ticket (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price float8,
    quantity integer,
    event_id integer references event(id)
);
CREATE TABLE booking (
    id SERIAL PRIMARY KEY,
    account_id integer references account(id),
    event_id integer references event(id),
    quantity integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account;
DROP TABLE event;
DROP TABLE ticket;
DROP TABLE booking;
-- +goose StatementEnd
