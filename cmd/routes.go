package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohit-ludbe/go-trivia/handlers"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.GetFacts)

	app.Post("/fact", handlers.CreateFact)

}
