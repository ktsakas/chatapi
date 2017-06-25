DROP TABLE contacts;
DROP TABLE channel_users;
DROP TABLE messages;
DROP TABLE channels;
DROP TABLE users;

-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    university VARCHAR(255),
    talking_to VARCHAR(255),
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
    is_group BOOLEAN,
    name VARCHAR(255),
    domain VARCHAR(255) UNIQUE,
    created_at TIMESTAMP
);

INSERT INTO channels (name, is_group, domain, created_at) VALUES
    ('Harvard University', true, 'harvard.edu', NOW()),
    ('Brown University', true, 'brown.edu', NOW()),
    ('Columbia University', true, 'columbia.edu', NOW()),
    ('Cornell University', true, 'cornell.edu', NOW()),
    ('Dartmouth College', true, 'dartmouth.edu', NOW()),
    ('University of Pennsylvania', true, 'upenn.edu', NOW()),
    ('Princeton University', true, 'princeton.edu', NOW()),
    ('Yale University', true, 'yale.edu', NOW());

-- Create channels to users table
CREATE TABLE channel_users (
    channel_id SERIAL REFERENCES channels (id),
    user_id SERIAL REFERENCES users (id),
    UNIQUE(channel_id, user_id)
);

-- Create messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    channel_id SERIAL REFERENCES channels (id),
    sender_id SERIAL REFERENCES users (id),
    created_at TIMESTAMP
);