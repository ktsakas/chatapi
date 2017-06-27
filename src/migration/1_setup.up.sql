-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY,
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
    usera_id UUID REFERENCES users (id),
    userb_id UUID REFERENCES users (id),
    UNIQUE(usera_id, userb_id)
);

-- Create channels table
CREATE TABLE channels (
    id UUID PRIMARY KEY,
    is_group BOOLEAN,
    name VARCHAR(255),
    domain VARCHAR(255) UNIQUE,
    created_at TIMESTAMP
);

INSERT INTO channels (id, name, is_group, domain, created_at) VALUES
    ('6b2d91a3-380b-4542-a0b6-269dd4948611', 'Harvard University', true, 'harvard.edu', NOW()),
    ('33c8a8c8-f7fa-48f4-bcbd-8d2dcde221ad', 'Brown University', true, 'brown.edu', NOW()),
    ('4be98e8b-c39e-4fd0-8a13-1f3b1368cf01', 'Columbia University', true, 'columbia.edu', NOW()),
    ('81c9a546-bb9d-467f-b6e0-6ba7eb52dd04', 'Cornell University', true, 'cornell.edu', NOW()),
    ('9e358765-f2f4-48b8-82e4-feca0cab1397', 'Dartmouth College', true, 'dartmouth.edu', NOW()),
    ('c1f611e3-9e1f-4d23-9b84-0be026db7be8', 'University of Pennsylvania', true, 'upenn.edu', NOW()),
    ('724c911c-b98e-441c-877c-d70cb42a8d61', 'Princeton University', true, 'princeton.edu', NOW()),
    ('d8176f7d-2f24-404c-ae10-43585c3a14ec', 'Yale University', true, 'yale.edu', NOW());

-- Create channels to users table
CREATE TABLE channel_users (
    channel_id UUID REFERENCES channels (id),
    user_id UUID REFERENCES users (id),
    UNIQUE(channel_id, user_id)
);

-- Create messages table
CREATE TABLE messages (
    id UUID PRIMARY KEY,
    channel_id UUID REFERENCES channels (id),
    sender_id UUID REFERENCES users (id),
    created_at TIMESTAMP
);