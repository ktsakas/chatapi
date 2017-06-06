CREATE DATABASE collegechat;
SET DATABASE = collegechat;

-- Create users table
CREATE TABLE account (
    uuid SERIAL PRIMARY KEY,
    email STRING UNIQUE,
    password STRING,
    sex STRING,
    university STRING,
    createdAt TIMESTAMP,
    updatedAt TIMESTAMP
);

-- Create contacts table (many2many: user <-> user)
CREATE TABLE contacts (
    accountA STRING,
    accountB STRING,
    UNIQUE(accountA, accountB)
);

-- Create messages table
CREATE TABLE messages (
    uuid SERIAL PRIMARY KEY,
    channel SERIAL,
    account SERIAL REFERENCES account (uuid)
);
