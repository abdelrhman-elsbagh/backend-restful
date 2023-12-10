package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

var dbConn *pgx.Conn

func InitDB() (*pgx.Conn, error) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	//conn, err := pgx.Connect(context.Background(), "postgres://postgres:Admin2030@localhost:5432/api_users")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	dbConn = conn

	return conn, nil
}

func CreateUsersTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            phone_number TEXT UNIQUE NOT NULL,
            otp TEXT NOT NULL,
            otp_expiration_time TIMESTAMP NULL
        );
    `)
	return err
}

func GetDB() *pgx.Conn {
	return dbConn
}

func CloseDB() {
	if dbConn != nil {
		dbConn.Close(context.Background())
	}
}
