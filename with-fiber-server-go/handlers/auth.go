package handlers

import (

	"github.com/gofiber/fiber/v2"
	"nested-comments/database"
	"nested-comments/utils"
)

func (h *Handlers) Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"message": "Invalid input",
				"error": err.Error(),
			},
		)
	}
	userPassw , err := h.Repo.GetUser(c.Context(), input.Email )
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	if userPassw == "" {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"message": "User not found",
			},
		)
	}
	if !utils.CheckPassword(userPassw, input.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"message": "Invalid password",
			},
		)
	}
	token, err := utils.GenerateToken("", input.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"token": token, "msg": "login success", "email": input.Email,
		},
	)
}


func (h *Handlers) Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"message": "Invalid input",
				"error": err.Error(),
			},
		)
	}

	hashedPw, err := utils.HashPassword(string(input.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	newUser, err := h.Repo.CreateUser(c.Context(), database.CreateUserParams{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hashedPw,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func (h *Handlers) Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*utils.Claims)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"user": user,
		},
	)
}