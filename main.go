package main

import (
	"gateway/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.GetConf()

	app := fiber.New(fiber.Config{
		AppName:   "GATEWAY",
		BodyLimit: 1024 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(func(c *fiber.Ctx) error {
		service := strings.Split(c.OriginalURL(), "/")[2]

		return c.Redirect(config.Routes[service] + c.OriginalURL())
	})

	app.Listen(":" + config.PORT)
}
