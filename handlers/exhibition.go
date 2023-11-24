package handlers

import (
	"ar-museum-backend/models"
	"ar-museum-backend/storage"
	"ar-museum-backend/utils"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"log"
	"net/url"
	"strconv"
	"time"
)

type ExhibitionQr struct {
	MuseumId uint64 `json:"museumId"`
	ExhibitionId uint64 `json:"exhibitionId"`
	Locale string `json:"locale"`
}

type ExhibitMetaInfo struct {
	Id uint64 `json:"id"`
	ModelUrl string `json:"model"`
	AudioUrl string `json:"audio"`
	DescriptionUrl string `json:"description"`
	ImageUrls []string `json:"images"`
}

func CreateLinks(prefix string, bucketName string) ([]string, error) {
	log.Print(prefix)
	var result []string
	ctx := context.Background()
	objectCh := storage.MIOClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Recursive: true,
		Prefix: prefix,
	})
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
		}
		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "attachment; filename=\"attach\"")
		url, err := storage.MIOClient.PresignedGetObject(ctx, bucketName, object.Key, time.Second * 300, reqParams)
		if err == nil {
			result = append(result, url.String())
		} else {
			log.Println(err.Error())
		}

	}

	if len(result) == 0 {
		return nil, errors.New("Empty bucket")
	}

	return result, nil
}

func GetExhibitMetaInfo(id uint64, prefix string, locale string) ExhibitMetaInfo {
	var result ExhibitMetaInfo
	prefixWithLocale := fmt.Sprintf("%s/%s", prefix, locale)
	/// get audio
	{
		audioUrls, err := CreateLinks(prefixWithLocale, "audio-bucket")
		if err != nil {
			log.Print(err.Error())
		} else {
			result.AudioUrl = audioUrls[0]
		}
	}
	/// get descriptions
	{
		descriptionUrls, err := CreateLinks(prefixWithLocale, "description-bucket")
		if err != nil {
			log.Print(err.Error())
		} else {
			result.DescriptionUrl = descriptionUrls[0]
		}
	}
	/// get model
	{
		modelUrls, err := CreateLinks(prefix, "models-bucket")
		if err != nil {
			log.Print(err.Error())
		} else {
			result.ModelUrl = modelUrls[0]
		}
	}

	/// get images
	{
		imagesUrls, err := CreateLinks(prefix, "photo-bucket")
		if err != nil {
			log.Print(err.Error())
		} else {
			result.ImageUrls = imagesUrls
		}
	}
	result.Id = id
	return result
}

func GetExhibitionInfo(fiberContext *fiber.Ctx) error {
	log.Print("test")
	var result []ExhibitMetaInfo
	exhibition := new(ExhibitionQr)
/*	if err := fiberContext.BodyParser(exhibition); err != nil {
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}*

 */
	i, _ := strconv.Atoi(fiberContext.Query("exhibitionId"))
	exhibition.ExhibitionId = uint64(i)
	exhibition.Locale = fiberContext.Query("locale")
	i, _ = strconv.Atoi(fiberContext.Query("museumId"))
	exhibition.MuseumId = uint64(i)
	log.Print("Struct ", exhibition)

	rows, _ := storage.DBCon.Raw("SELECT exhibits.id FROM exhibits WHERE exhibits.client_id = ? AND exhibits.exhibition_id = ?", exhibition.MuseumId, exhibition.ExhibitionId).Rows()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		prefix := fmt.Sprintf("%d/%d/%d", exhibition.MuseumId, exhibition.ExhibitionId, id)
		result = append(result, GetExhibitMetaInfo(uint64(id), prefix, exhibition.Locale))
	}

	return fiberContext.Status(200).JSON(result)
}

func CreateExhibition(fiberContext *fiber.Ctx) error {
	exhibition := new(models.Exhibition)
	if err := fiberContext.BodyParser(exhibition); err != nil {
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	storage.DBCon.Create(&exhibition)
	qrSource := fmt.Sprintf("%s/%s", strconv.FormatUint(exhibition.ClientId, 10),
		strconv.FormatUint(exhibition.ID, 10))
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

	return fiberContext.Status(200).JSON(exhibition)
}

func GetExhibitions(context *fiber.Ctx) error {
	var exhibitions []models.Exhibition
	storage.DBCon.Find(&exhibitions)
	return context.Status(200).JSON(exhibitions)
}
