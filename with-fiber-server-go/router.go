package main

import (
	"nested-comments/handlers"
	"nested-comments/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// protected rotues middlware, pass user in context if valid
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"message": "Missing token"},
			)
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{"message": "Invalid token",
					"error": err.Error(),
				},
			)
		}
		c.Locals("user", claims)
		return c.Next()
	}
}

func SetupRoutes(app *fiber.App, handlers *handlers.Handlers) {
	api := app.Group("/v1", logger.New())
	api.Get("/test", handlers.Hello)

	api.Get("/", handlers.GetPostFeed)
	api.Post("/", Protected(), handlers.CreatePost)

	api.Get("/:postId", handlers.GetPost)

	api.Get("/:postId/comments", handlers.GetComments)
	api.Post("/:postId/comments", Protected(), handlers.CreateComment)
	api.Post("/:postId/comments/:parentId", Protected(), handlers.CreateComment)

	auth := api.Group("/auth")
	auth.Get("/me", Protected(), handlers.Me)
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)

}
