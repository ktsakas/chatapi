DROP TABLE contacts;
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
    name VARCHAR(255) UNIQUE,
    domain VARCHAR(255)
);

INSERT INTO channels (name, domain) VALUES
    ('Harvard University', 'harvard.edu'),
    ('Brown University', 'brown.edu'),
    ('Columbia University', 'columbia.edu'),
    ('Cornell University', 'cornell.edu'),
    ('Dartmouth College', 'dartmouth.edu'),
    ('University of Pennsylvania', 'upenn.edu'),
    ('Princeton University', 'princeton.edu'),
    ('Yale University', 'yale.edu');

-- Create messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    channel_id SERIAL REFERENCES channels (id),
    sender_id SERIAL REFERENCES users (id)
);
