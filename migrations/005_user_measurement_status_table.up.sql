CREATE TABLE IF NOT EXISTS user_measurement_status (
    id CHAR(26) PRIMARY KEY,
    user_id CHAR(26) NOT NULL,
    measurement_date_id CHAR(26) NOT NULL,
    image_data_id CHAR(26) DEFAULT NULL,
    has_registered BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (measurement_date_id) REFERENCES measurement_date(id),
    FOREIGN KEY (image_data_id) REFERENCES image_data(id)
);