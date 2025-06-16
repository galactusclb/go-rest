package product

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"CLB-go-rest/validation"
)

type (
	Product struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		CategoryID string `json:"category_id"`
	}
)

var products = []Product{
	{ID: "1", Name: "Smartphone", CategoryID: "1"},
	{ID: "2", Name: "Laptop", CategoryID: "1"},
}

func GetProducts(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, item := range products {
		if fmt.Sprint(item.ID) == id {
			return c.Status(http.StatusOK).JSON(item)
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct Product
	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Body"})
	}

	productValidation := validation.NewValidator()
	errors := productValidation.Validate(newProduct)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	newProduct.ID = fmt.Sprintf("%d", len(products)+1)
	products = append(products, newProduct)

	return c.Status(fiber.StatusCreated).JSON(newProduct)
}
