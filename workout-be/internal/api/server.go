package api

import (
	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config config.Config
	store  db.Store
	app    *fiber.App
}

func NewServer(config config.Config, store db.Store) *Server {
	return &Server{
		config: config,
		store:  store,
		app:    fiber.New(),
	}
}

func (s *Server) SetupRoutes() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	userController := controllers.NewUserController(s.store)

	userRoutes := s.app.Group("/users")
	userRoutes.Post("/", userController.CreateUser)
}

func (s *Server) Start(address string) error {
	s.SetupRoutes()
	return s.app.Listen(address)
}
