CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(60) NOT NULL,
    name VARCHAR(60),
    image_uri TEXT,
    preference PREFERENCE_TYPE,
    weight_unit WEIGHT_UNIT,
    height_unit HEIGHT_UNIT,
    weight DECIMAL(5,2),
    height DECIMAL(5,2),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT unique_email UNIQUE (email),
    CONSTRAINT weight_range CHECK (weight >= 10 AND weight <= 1000),
    CONSTRAINT height_range CHECK (height >= 3 AND height <= 250)
);

CREATE TABLE activities (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    activity_type ACTIVITY_TYPE NOT NULL,
    done_at TIMESTAMP NOT NULL,
    duration_minutes INTEGER NOT NULL,
    calories_burned DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT duration_positive CHECK (duration_minutes > 0)
);

CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    original_name VARCHAR(255) NOT NULL,
    content_type VARCHAR(100) NOT NULL,
    size INTEGER NOT NULL,
    uri TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT file_size_limit CHECK (size <= 102400)
);