CREATE TABLE booking_types (
    id SERIAL PRIMARY KEY,
    booking_type VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO booking_types (booking_type, description) VALUES
('Car Only', 'Rent Car only'),
('Car & Driver', 'Rent Car and a Driver');
