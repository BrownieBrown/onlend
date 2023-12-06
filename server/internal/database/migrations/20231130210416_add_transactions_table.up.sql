CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    sender_id UUID REFERENCES users(id),
    receiver_id UUID REFERENCES users(id),
    amount NUMERIC(19,4),
    transaction_type INT,
    status INT
);
