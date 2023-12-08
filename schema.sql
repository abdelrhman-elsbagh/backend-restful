-- create_users_table.sql

-- This file contains the SQL statement to create the 'users' table.
-- name: CreateTable :one
CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL,
       phone_number TEXT UNIQUE NOT NULL,
       otp TEXT NOT NULL,
       otp_expiration_time TIMESTAMP NULL
);