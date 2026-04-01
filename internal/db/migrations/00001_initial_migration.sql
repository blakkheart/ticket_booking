-- +goose Up
-- +goose StatementBegin
CREATE TABLE account (
    id BIGSERIAL PRIMARY KEY,
    
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    role VARCHAR(128) NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE event (
    id BIGSERIAL PRIMARY KEY,
    
    title VARCHAR(255) NOT NULL,
    description TEXT,
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ,
    location VARCHAR(255) NOT NULL,
    status VARCHAR(128) NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE ticket_type (
    id BIGSERIAL PRIMARY KEY,
    
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price NUMERIC(10,2) NOT NULL,
    available_quantity integer NOT NULL,
    
    event_id BIGINT NOT NULL references event(id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE booking (
    id BIGSERIAL PRIMARY KEY,
    
    account_id BIGINT NOT NULL references account(id) ON DELETE CASCADE,
    status VARCHAR(128) NOT NULL,

    expires_at TIMESTAMPTZ,
    paid_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_booking_account_id ON booking(account_id);
CREATE INDEX idx_booking_status ON booking(status);
CREATE INDEX idx_booking_expires_at ON booking(expires_at);

CREATE TABLE booking_items (
    id BIGSERIAL PRIMARY KEY,

    booking_id BIGINT NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    ticket_type_id BIGINT NOT NULL REFERENCES ticket_type(id) ON DELETE CASCADE,

    quantity INTEGER NOT NULL,
    price_at_booking NUMERIC(10,2) NOT NULL
);
CREATE INDEX idx_booking_items_booking_id ON booking_items(booking_id);
CREATE INDEX idx_booking_items_ticket_type_id ON booking_items(ticket_type_id);

CREATE TABLE refresh_token (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL,
    token_hash VARCHAR(512) NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked BOOLEAN NOT NULL DEFAULT FALSE, 
    replaced_by BIGINT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    FOREIGN KEY (user_id) REFERENCES account(id),
    FOREIGN KEY (replaced_by) REFERENCES refresh_token(id)
);
CREATE INDEX idx_refresh_token_user_id ON refresh_token(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account CASCADE;
DROP TABLE IF EXISTS event CASCADE;
DROP TABLE IF EXISTS ticket_type CASCADE;
DROP TABLE IF EXISTS booking CASCADE;
DROP TABLE IF EXISTS booking_items CASCADE;
DROP TABLE IF EXISTS refresh_token CASCADE;
-- +goose StatementEnd
