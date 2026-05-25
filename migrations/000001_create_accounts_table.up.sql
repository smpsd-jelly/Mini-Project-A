CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY
    account_number VARCHAR(20) UNIQUE NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    citizen_id VARCHAR(13) UNIQUE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    account_type VARCHAR(20) NOT NULL
    CHECK (account_type IN ('SAVING', 'CURRENT'))
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',
    CHECK (status IN ('ACTIVE','CLOSED')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);