package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohit-ludbe/go-trivia/database"
	"github.com/rohit-ludbe/go-trivia/models"
)

func GetFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
