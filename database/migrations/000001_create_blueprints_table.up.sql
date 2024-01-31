CREATE TABLE blueprints (
    id VARCHAR(32) PRIMARY KEY,
    version VARCHAR(255) NOT NULL,
    brand_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
