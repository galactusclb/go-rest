package product

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	router.Get("", GetProducts)
	router.Get(":id", GetProduct)
	router.Post("", CreateProduct)
}
