CREATE DATABASE collegechat;
SET DATABASE = collegechat;

-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    sex VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Create contacts table (many2many: user <-> user)
CREATE TABLE contacts (
    usera_id SERIAL REFERENCES users (id),
    userb_id SERIAL REFERENCES users (id),
    UNIQUE(usera_id, userb_id)
);

-- Create channels table
CREATE TABLE channels (
    id SERIAL PRIMARY KEY,
    organization VARCHAR(255) UNIQUE
);

-- Create messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    channel_id SERIAL REFERENCES channels (id),
    sender_id SERIAL REFERENCES users (id)
);
