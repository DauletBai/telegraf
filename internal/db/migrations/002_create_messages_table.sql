CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    sender_id INT REFERENCE users(id),
    receiver_id INT REFERENCE users(id),
    content TEXT NOT NULL,
    message_type VARCHAR(20) NOT NULL, -- text, audio, video
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
