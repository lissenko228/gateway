package router

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	api := app.Group("/api")

	UsersRouter(api)
}
