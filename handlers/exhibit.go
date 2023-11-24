package handlers

import (
	"ar-museum-backend/models"
	"ar-museum-backend/storage"
	"ar-museum-backend/utils"
	"bytes"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"strconv"
)

type Exhibit struct {
	ID uint64 `json:"id"`
	ExhibitionId uint64 `json:"exhibitionId"`
	ClientId uint64 `json:"clientId"`
}

func CreateExhibit(fiberContext *fiber.Ctx) error {
	exhibit := new(Exhibit)
	if err := fiberContext.BodyParser(exhibit); err != nil {
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dbExhibit := models.Exhibit {
		Name: "Test",
		ExhibitionId: exhibit.ExhibitionId,
		ClientId: exhibit.ClientId,
	}

	storage.DBCon.Create(&dbExhibit)
	qrSource := fmt.Sprintf("%s/%s/%s", strconv.FormatUint(exhibit.ClientId, 10),
		strconv.FormatUint(exhibit.ExhibitionId, 10), strconv.FormatUint(dbExhibit.ID, 10))
	png, err := utils.GenerateQrCode(qrSource)
	if err != nil {
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	contentType := "application/octet-stream"
	reader := bytes.NewReader(png)
	ctx := context.Background()
	_, err = storage.MIOClient.PutObject(ctx, "qr-bucket", qrSource, reader, reader.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return fiberContext.Status(200).JSON(exhibit)
}

func GetExhibits(context *fiber.Ctx) error {
	var exhibits []models.Exhibit
	storage.DBCon.Find(&exhibits)
	return context.Status(200).JSON(exhibits)
}