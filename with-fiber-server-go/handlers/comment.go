package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"nested-comments/database"
	"nested-comments/utils"
	"strconv"
)

func (h *Handlers) CreateComment(c *fiber.Ctx) error {
	type NewCommentInput struct {
		Body string `json:"body"`
	}
	postId, err := strconv.Atoi(c.Params("postId"))
	parentId, err2 := strconv.Atoi(c.Params("parentId"))
	if err != nil || len(c.Params("parentId")) > 0 &&  err2 != nil {
		return  c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"message": "Invalid input",
				"error": err.Error(),
			},)
	}


	var input NewCommentInput
	if err = c.BodyParser(&input); err != nil  {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"message": "Invalid input",
				"error": err.Error(),
			},
		)
	}

	user := c.Locals("user").(*utils.Claims)
	var newCommentArgs = database.CreateCommentParams{
		Content:   input.Body,
		UserEmail: user.Email,
		PostID:    int32(postId),
		ParentID:   sql.NullInt32{Valid: parentId > 0, Int32: int32(parentId)},
	}
	comment, err := h.Repo.CreateComment(c.Context(), newCommentArgs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error",
				"error": err.Error(),
			},
		)
	}
	//success
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{"message": "Comment created", "comment": fiber.Map{
			"id":        comment.ID,
			"body":      comment.Content,
			"postId":    comment.PostID,
			"userEmail": comment.UserEmail,
			"parentId":  comment.ParentID.Int32,
		},
		},
	)
}
