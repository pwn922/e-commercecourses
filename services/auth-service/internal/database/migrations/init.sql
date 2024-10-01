CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE DATABASE auth_service;

\c auth_service;

CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

INSERT INTO roles (role_name, description) VALUES
('student', 'Role for students'),
('admin', 'Role for administrators'),
('instructor', 'Role for instructors')
ON CONFLICT (role_name) DO NOTHING;