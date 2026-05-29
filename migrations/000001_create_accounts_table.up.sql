CREATE SEQUENCE account_number_seq START 1;
CREATE SEQUENCE IF NOT EXISTS account_number_seq
START WITH 1
INCREMENT BY 1;

CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    citizen_id VARCHAR(13) UNIQUE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    account_type VARCHAR(20) NOT NULL CHECK (account_type IN ('SAVING', 'CURRENT')),
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0),
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN ('ACTIVE','CLOSED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

CREATE OR REPLACE FUNCTION generate_account_number()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.account_number IS NULL OR NEW.account_number = '' THEN
        NEW.account_number := LPAD(nextval('account_number_seq')::text, 10, '0');
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_accounts_generate_account_number
BEFORE INSERT ON accounts
FOR EACH ROW
EXECUTE FUNCTION generate_account_number();