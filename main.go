package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/shop/greeting", func(c *fiber.Ctx) error {
		response := "Hello World"
		log.Println("Response:", response)
		return c.SendString("Hello World")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
