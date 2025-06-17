package utils

import "github.com/gofiber/fiber/v2"

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
