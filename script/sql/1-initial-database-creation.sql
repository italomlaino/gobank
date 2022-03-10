START TRANSACTION;

-- Database
CREATE DATABASE IF NOT EXISTS gobank;
USE gobank;

-- Account
CREATE TABLE account
(
    id              INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    document_number INT UNSIGNED NOT NULL
);

-- Transaction
CREATE TABLE account_transaction
(
    id                INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    account_id        INT UNSIGNED NOT NULL,
    operation_type_id INT UNSIGNED NOT NULL,
    amount            INT          NOT NULL,
    event_data        TIMESTAMP    NOT NULL,
    CONSTRAINT fk_account_transaction_account FOREIGN KEY (account_id) REFERENCES account (id)
);


COMMIT;