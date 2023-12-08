package main

import (
	"awesomeProject/database"
	"awesomeProject/db"
	"awesomeProject/routes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	conn, err := database.InitDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing database: %v\n", err)
		os.Exit(1)
	}
	defer database.CloseDB()

	q := db.New(conn)

	err = database.CreateUsersTable(context.Background(), conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.Use(routes.SetupAPIHandler(q))

	routes.SetupRoutes(r)

	r.Run(":8080")
}
