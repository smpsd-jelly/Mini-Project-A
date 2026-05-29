DROP TRIGGER IF EXISTS before_insert_accounts_generate_account_number ON accounts;

DROP FUNCTION IF EXISTS generate_account_number();

DROP TABLE IF EXISTS accounts;

DROP SEQUENCE IF EXISTS account_number_seq;