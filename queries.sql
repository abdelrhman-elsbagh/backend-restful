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

-- name: CreateUser :one
INSERT INTO users (
    name, phone_number, otp, otp_expiration_time
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: UpdateUserOTP :exec
UPDATE users SET
    otp = $2,
    otp_expiration_time = $3
WHERE id = $1;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users
WHERE phone_number = $1 LIMIT 1;