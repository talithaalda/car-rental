CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    car_id INT REFERENCES cars(id),
    driver_id INT REFERENCES drivers(id),
    book_type_id INT REFERENCES booking_types(id),
    start_rent TIMESTAMP NOT NULL,
    end_rent TIMESTAMP NOT NULL,
    total_cost INT NOT NULL,
    discount INT,
    total_driver_cost INT,
    finished BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);