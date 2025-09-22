-- +goose Up
-- Insert payment methods
INSERT INTO payment_methods (name, code)
VALUES 
  ('Credit Card', 'CREDIT'),
  ('Debit Card', 'DEBIT'),
  ('PIX', 'PIX');

-- Insert clients
INSERT INTO clients (msisdn, balance, credit_limit)
VALUES 
  ('005', 150.75, 500.00),
  ('5521988888888', 50.00, 200.00);

-- Insert offers
INSERT INTO offers (name, value)
VALUES
  ('Starter Pack', 9.99),
  ('Premium Pack', 29.99);

-- Insert credit cards
INSERT INTO credit_cards (card_number, brand)
VALUES
  ('4111111111111111', 'VISA'),
  ('5500000000000004', 'MASTERCARD');

-- Link clients to credit cards using IDs
INSERT INTO user_credit_cards (client_id, card_id)
SELECT c.id, cc.id
FROM clients c
JOIN credit_cards cc
ON (c.msisdn = '005' AND cc.card_number = '4111111111111111')
   OR (c.msisdn = '5521988888888' AND cc.card_number = '5500000000000004');

-- Insert orders with correct client and payment method IDs
INSERT INTO orders (order_number, client_id, amount, payment_method_id)
SELECT 'ORD-1001', c.id, 29.99, pm.id
FROM clients c, payment_methods pm
WHERE c.msisdn = '005' AND pm.code = 'CREDIT'
UNION ALL
SELECT 'ORD-1002', c.id, 9.99, pm.id
FROM clients c, payment_methods pm
WHERE c.msisdn = '5521988888888' AND pm.code = 'PIX';