CREATE TABLE memberships (
    id SERIAL PRIMARY KEY,
    membership_name VARCHAR(255) NOT NULL,
    discount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert data
INSERT INTO memberships (membership_name, discount) VALUES
('Bronze', 4),
('Silver', 7),
('Gold', 15);


