package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"CLB-go-rest/location"
	"CLB-go-rest/product"
)

type (
	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Welcome to Product API")
	})

	v1 := app.Group("/api/v1")

	product.RegisterRoutes(v1.Group("/products"))
	location.RegisterRoutes((v1.Group("/locations")))

	PORT := os.Getenv("PORT")
	log.Printf("Listening on port : %s", PORT)
	log.Fatal(app.Listen(":" + PORT))
}
