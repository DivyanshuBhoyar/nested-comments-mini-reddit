package handlers

import (
	"github.com/gofiber/fiber/v2"
	"nested-comments/database"
)

type Handlers struct {
	Repo *database.Repo
}

func NewHandlers(repo *database.Repo) *Handlers {
	return &Handlers{
		Repo: repo,
	}
}

func (h *Handlers) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
