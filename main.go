package main

import (
	"ar-museum-backend/handlers"
	"ar-museum-backend/storage"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/create_client", handlers.CreateClient)
	app.Get("/get_clients", handlers.GetClients)
	app.Post("/create_exhibition", handlers.CreateExhibition)
	app.Get("/get_exhibitions", handlers.GetExhibitions)
	app.Get("/get_exhibitionqr", handlers.GetClients)
	app.Get("/test", handlers.GetExhibitionInfo)
	app.Post("/create_exhibit", handlers.CreateExhibit)
	app.Get("/get_exhibits", handlers.GetExhibits)
}

func main() {
	postgreConfig := &storage.PostgreConfig{
		Host: "db",
		Port: "5432",
		User: "admin",
		Password: "12345678",
		DBName: "armuseum_db",
		SSL: "disable",
	}
	var err error
	storage.DBCon, err = storage.CreateConnection(postgreConfig)
	if err != nil {
		log.Fatal("Failed to connect to db\n");
		os.Exit(2)
	}
	log.Println("Connected to postgres")

	mioConfig := &storage.MioConfig{
		//Host: "s3service:9000",
		Host: "188.232.151.86:85",
		AccessKey: "Yc0BlzTSb8okc3qKgKMH",
		SecretKey: "iZ6N3m1Z2uGABJ8umwlZ5JJmEVQBaEK56gHNWNZm",
		Secure: false,
	}

	storage.MIOClient, err = storage.CreateMIOClient(mioConfig)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(storage.MIOClient)
	/*
	objectName := "testqr.png"
	contentType := "application/octet-stream"

	var png []byte
	png, _ = qrcode.Encode("https://example.org", qrcode.Medium, 256)
	reader := bytes.NewReader(png)

	ctx := context.Background()
	info, err := minioClient.PutObject(ctx, "test-bucket", objectName, reader, reader.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)*/

	app := fiber.New()
	SetupRoutes(app)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world")
	})
	app.Listen(":3000")
}
