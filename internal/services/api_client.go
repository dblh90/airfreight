package services

import (
	"encoding/csv"
	"encoding/json"
	"github.com/dbl90/airfreight/internal/models"
	"github.com/dbl90/airfreight/internal/models/config"
	"github.com/dbl90/airfreight/internal/models/db"
	"github.com/go-playground/validator/v10"
	"github.com/gocarina/gocsv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"io"
	"log"
)

type APIClient struct {
	config *config.AppConfig
	db     *db.DBClient
	App    *fiber.App
}

func NewAPIClient(config *config.AppConfig, client *db.DBClient) *APIClient {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               "airfreight app",
	})
	app.Use(logger.New(logger.Config{
		Format:   "${pid} ${time} ${locals:requestid} ${status} - ${method} ${path}​\n",
		TimeZone: "UTC",
	}))

	api := app.Group("/api").Name("api")
	api.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"hamza": "123456",
			"admin": "aabbcc",
		},
		Authorizer:   Authorizer,
		Unauthorized: Unauthorizer,
	}))
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
		})
	})
	api.Post("/air_export/mawbs/:mawbsNum", func(ctx *fiber.Ctx) error {
		return createMawb(ctx, client)
	})
	api.Post("/air_export/shipments/:hawbNum", func(ctx *fiber.Ctx) error {
		return createHawb(ctx, client)
	})
	api.Post("/air_export/mawbs/:mawbsNum/shipments/attach/:hawbNum", func(ctx *fiber.Ctx) error {
		return attach(ctx, client)
	})
	return &APIClient{
		config: config,
		db:     client,
		App:    app,
	}
}

func attach(c *fiber.Ctx, client *db.DBClient) error {
	form, err := c.FormFile("attachment")
	file, err := form.Open()
	defer file.Close()
	if err != nil {
		return err
	}

	gocsv.SetCSVReader(func(reader io.Reader) gocsv.CSVReader {
		r := csv.NewReader(reader)
		r.Comma = ';'
		return r
	})

	var shipments []*models.Shipment
	if err := gocsv.UnmarshalMultipartFile(&file, &shipments); err != nil {
		panic(err)
	}
	var hawbs = make([]models.Hawb, 0)
	for _, shipment := range shipments {
		mawbByNumber := client.GetMawbByNumber(shipment.Mawb)
		if mawbByNumber == nil {
			log.Println("Mawb not found, creating new one")
			mawbByNumber := models.Mawb{Number: shipment.Mawb}
			client.Insert(&mawbByNumber)
			log.Println("Mawb created", mawbByNumber.ID)
		}

		err := validator.New().Struct(shipment)
		if err != nil {
			log.Println("Error validating shipment, skipping", shipment, err)
			continue
		}

		hawbs = append(hawbs, models.Hawb{
			Origin:      shipment.Origin,
			Destination: shipment.Destination,
			Consignor:   shipment.Consignee,
			Consignee:   shipment.Consignee,
			Content:     shipment.Content,
			Pieces:      shipment.Pieces,
			Number:      shipment.Hawb,
			MawbID:      mawbByNumber.ID,
			Weight:      shipment.Weight,
		})
	}
	client.Insert(&hawbs)
	log.Println("Request received")
	return c.JSON(fiber.Map{"message": "Attachment uploaded successfully", "file_size": len(shipments)})
}

func createHawb(ctx *fiber.Ctx, client *db.DBClient) error {
	hawbb := &models.Hawb{}
	err := json.Unmarshal(ctx.Body(), hawbb)
	if err != nil {
		return err
	}
	hawbNum := ctx.Params("hawbNum")
	hawbb.Number = hawbNum
	result := client.Insert(hawbb)
	if result.Error != nil {
		response := ctx.Response()
		response.SetBody([]byte("Error inserting hawb"))
		response.SetStatusCode(fiber.StatusInternalServerError)
	}
	return ctx.JSON(fiber.Map{"hawb": hawbb})
}

func createMawb(ctx *fiber.Ctx, client *db.DBClient) error {
	mawb := &models.Mawb{}
	err := json.Unmarshal(ctx.Body(), mawb)
	if err != nil {
		return err
	}
	mawbsNum := ctx.Params("mawbsNum")
	mawb.Number = mawbsNum
	insert := client.Insert(mawb)
	if insert.Error != nil {
		response := ctx.Response()
		response.SetBody([]byte("Error inserting mawb"))
		response.SetStatusCode(fiber.StatusInternalServerError)
	}
	return ctx.JSON(fiber.Map{"mawb": mawb})
}
