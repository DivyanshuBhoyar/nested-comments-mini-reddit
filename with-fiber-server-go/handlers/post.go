package handlers

import (
	"github.com/gofiber/fiber/v2"
	"nested-comments/database"
	"nested-comments/utils"
	"strconv"
)

func (h *Handlers) CreatePost(c *fiber.Ctx) error {
	type NewPostInput struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	var input NewPostInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"message": "Invalid input",
				"error": err.Error(),
			},
		)
	}
	user := c.Locals("user").(*utils.Claims)
	var newPostArgs = database.CreatePostParams{
		Title:       input.Title,
		Body:        input.Body,
		AuthorEmail: user.Email,
	}
	post, err := h.Repo.CreatePost(c.Context(), newPostArgs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{"message": "Post created", "post": fiber.Map{
			"id":    post.ID,
			"body":  post.Body,
			"title": post.Title,
		},
		},
	)
}

func (h *Handlers) GetPostFeed(c *fiber.Ctx) error {
	posts, err := h.Repo.GetPostFeed(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"message": "Posts feed, todo: pagination", "posts": posts},
	)
}

func (h *Handlers) GetPost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	id, err := strconv.Atoi(postId)
	post, err := h.Repo.GetPost(c.Context(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"message": "Post", "post": post},
	)
}


func (h *Handlers) GetComments(c *fiber.Ctx) error {
	postId := c.Params("postId")
	id, err := strconv.Atoi(postId)
	comments, err := h.Repo.GetComments(c.Context(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"message": "Comments", "comments": comments},
	)
}