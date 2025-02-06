CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    car_id INT REFERENCES cars(id),
    total_cost INT NOT NULL,
    start_rent TIMESTAMP NOT NULL,
    end_rent TIMESTAMP NOT NULL,
    finished BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);