CREATE TABLE users (
                       guid      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       login      VARCHAR(255) NOT NULL,
                       password   VARCHAR(255) NOT NULL,
                       name       VARCHAR(255) NOT NULL,
                       email      VARCHAR(255) NOT NULL,
                       verify_code VARCHAR(255),
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ DEFAULT NOW()
);