-- +goose Up
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    msisdn VARCHAR(20) UNIQUE NOT NULL,
    balance NUMERIC(12,2) DEFAULT 0,
    credit_limit NUMERIC(12,2) DEFAULT 0
);

CREATE TABLE credit_cards (
    id SERIAL PRIMARY KEY,
    card_number VARCHAR(20) UNIQUE NOT NULL,
    brand VARCHAR(50) NOT NULL
);

CREATE TABLE user_credit_cards (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    card_id INT NOT NULL REFERENCES credit_cards(id) ON DELETE CASCADE,
    UNIQUE (client_id, card_id)
);

CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE offers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    value NUMERIC(12,2) NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number VARCHAR(100) UNIQUE NOT NULL,
    client_id INT NOT NULL REFERENCES clients(id),
    offer_id INT REFERENCES offers(id),
    amount NUMERIC(12,2) NOT NULL,
    payment_method_id INT NOT NULL REFERENCES payment_methods(id),
    created_at TIMESTAMP DEFAULT NOW()
);
