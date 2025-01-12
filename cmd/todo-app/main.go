package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	time.Sleep(time.Second)
	log.Fatal(app.Listen(":3000"))
}
