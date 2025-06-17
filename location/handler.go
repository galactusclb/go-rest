package location

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	Location struct {
		ID          int64   `json:"id"`
		Name        string  `json:"number"`
		Description string  `json:"description"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	}
)

func GetLocations(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	limit := c.QueryInt("limit")

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 3
	}

	url := fmt.Sprintf("https://api.escuelajs.co/api/v1/locations?page=%d&size=%d", page, limit)

	response, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch data",
		})
	}

	defer response.Body.Close()

	var data []Location
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode response",
		})
	}

	responseObj := GeneratePaginationObj(data, page, limit)

	return c.Status(fiber.StatusOK).JSON(responseObj)
}

func GeneratePaginationObj[T any](data []T, page int, limit int) fiber.Map {
	nextPage := page + 1

	responseObj := fiber.Map{
		"data": &data,
		"meta": fiber.Map{
			"page":     page,
			"limit":    limit,
			"nextPage": nextPage,
		},
	}

	if page > 1 {
		if meta, ok := responseObj["meta"].(fiber.Map); ok {
			meta["prevPage"] = page - 1
		}
	}

	return responseObj
}
