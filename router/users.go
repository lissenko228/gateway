package router

import "github.com/gofiber/fiber/v2"

func UsersRouter(router fiber.Router) {
	users := router.Group("/users")

	users.Get("/*", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "OK"})
	})
}
