CREATE TABLE driver_incentives (
    id SERIAL PRIMARY KEY,
    booking_id INT NOT NULL,
    incentive INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
