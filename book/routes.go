package book

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	router.Get("", GetBooks)
	router.Post("", CreateBook)
}
