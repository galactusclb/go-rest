package location

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"CLB-go-rest/utils"
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

	responseObj := utils.GeneratePaginationObj(data, page, limit)

	return c.Status(fiber.StatusOK).JSON(responseObj)
}
