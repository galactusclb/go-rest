package location

import (
	"encoding/json"
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

	response, err := http.Get("https://api.escuelajs.co/api/v1/locations")
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

	return c.Status(fiber.StatusOK).JSON(data)
}
