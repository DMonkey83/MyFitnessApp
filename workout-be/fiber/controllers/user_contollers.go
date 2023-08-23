package controllers

import (
	"context"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	store *db.Store
}

func NewUserController(store db.Store) *UserController {
	return &UserController{store: store}
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user db.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdUser, err := uc.store.CreateUser(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to crate user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
