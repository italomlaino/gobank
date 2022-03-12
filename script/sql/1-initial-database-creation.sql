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

-- Operation Types
CREATE TABLE operation_types
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    description_pt VARCHAR(50)  NOT NULL
);

INSERT INTO operation_types (id, description_pt)
VALUES (1, 'COMPRA A VISTA'),
       (2, 'COMPRA PARCELADA'),
       (3, 'SAQUE'),
       (4, 'PAGAMENTO');

-- Transaction
CREATE TABLE transactions
(
    id                INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    account_id        INT UNSIGNED NOT NULL,
    operation_type_id INT UNSIGNED NOT NULL,
    amount            INT          NOT NULL,
    event_data        TIMESTAMP    NOT NULL,
    CONSTRAINT fk_transactions_accounts FOREIGN KEY (account_id) REFERENCES accounts (id),
    CONSTRAINT fk_operation_types FOREIGN KEY (operation_type_id) REFERENCES operation_types (id)
);
CREATE INDEX idx_transactions_account_id ON transactions (account_id);
CREATE INDEX idx_transactions_operation_type_id ON transactions (operation_type_id);

COMMIT;