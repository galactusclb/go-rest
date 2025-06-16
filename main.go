package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"CLB-go-rest/product"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Welcome to Product API")
	})

	v1 := app.Group("/api/v1")

	productRoute := v1.Group("/products")
	productRoute.Get("", product.GetProducts)
	productRoute.Get(":id", product.GetProduct)

	PORT := os.Getenv("PORT")
	log.Printf("Listening on port : %s", PORT)
	log.Fatal(app.Listen(":" + PORT))
}
