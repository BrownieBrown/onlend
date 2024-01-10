CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    account_type VARCHAR,
    balance NUMERIC(19,4)
);
