package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"nested-comments/database"
	"nested-comments/handlers"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	conn, err := sql.Open("postgres", fmt.Sprintf(
		"dbname=%s password=%s user=%s sslmode=disable", "nested_comments", "change_me", "postgres"))
	if err != nil {
		panic(err)
	}
	repo := database.NewRepo(conn)
	handlers := handlers.NewHandlers(repo)

	defer conn.Close()

	SetupRoutes(app, handlers)

	app.Listen(":3000")
}
