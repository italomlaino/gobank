START TRANSACTION;

-- Database
CREATE DATABASE IF NOT EXISTS gobank;
USE gobank;

-- Account
CREATE TABLE accounts
(
    id              INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    document_number INT UNSIGNED NOT NULL
);
CREATE INDEX idx_accounts_document_number ON accounts (document_number);

-- Transaction
CREATE TABLE transactions
(
    id                INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    account_id        INT UNSIGNED NOT NULL,
    operation_type_id INT UNSIGNED NOT NULL,
    amount            INT          NOT NULL,
    event_data        TIMESTAMP    NOT NULL,
    CONSTRAINT fk_transactions_accounts FOREIGN KEY (account_id) REFERENCES accounts (id)
);
CREATE INDEX idx_transactions_account_id ON transactions (account_id);


COMMIT;