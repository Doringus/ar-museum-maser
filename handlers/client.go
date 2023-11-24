package handlers

import (
	"ar-museum-backend/models"
	"ar-museum-backend/storage"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateClient(context *fiber.Ctx) error {
	client := new(models.Client)
	if err := context.BodyParser(client); err != nil {
		log.Println("Error")
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	storage.DBCon.Create(&client)
	return context.Status(200).JSON(client)
}

func GetClients(context *fiber.Ctx) error {
	var clients []models.Client
	storage.DBCon.Find(&clients)
	return context.Status(200).JSON(clients)
}
