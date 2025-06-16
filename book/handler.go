package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"CLB-go-rest/utils"
)

type (
	Book struct {
		Title       string   `json:"title"`
		Price       int      `json:"price"`
		Description string   `json:"description"`
		CategoryId  int      `json:"categoryId"`
		Images      []string `json:"images"`
	}

	BookResponse struct {
		Title       string   `json:"title"`
		Slug        string   `json:"slug"`
		Price       int      `json:"price"`
		Description string   `json:"description"`
		Images      []string `json:"images"`
		Category    struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Image string `json:"image"`
			Slug  string `json:"slug"`
		} `json:"category"`
		ID         int       `json:"id"`
		CreationAt time.Time `json:"creationAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	}
)

var API_URL = "https://api.escuelajs.co/api/v1"

func GetBooks(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	limit := c.QueryInt("limit")

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 3
	}

	url := fmt.Sprintf("%s/products?page=%d&size=%d", API_URL, page, limit)

	response, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch data",
		})
	}

	defer response.Body.Close()

	var data []BookResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode response",
		})
	}

	responseObj := utils.GeneratePaginationObj[BookResponse](data, page, limit)

	return c.Status(fiber.StatusOK).JSON(responseObj)
}

func CreateBook(c *fiber.Ctx) error {
	var newBook Book

	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	bodyBytes, err := json.Marshal(newBook)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to encode request body",
		})
	}

	url := fmt.Sprintf("%s/products", API_URL)
	resp, err := http.Post(url, "application/json",
		bytes.NewReader(bodyBytes))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send POST request",
		})
	}
	defer resp.Body.Close()

	var created BookResponse
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode response",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(created)
}
