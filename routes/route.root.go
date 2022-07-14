package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
)

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			webResponse := helpers.APIResponse(fiber.StatusNotFound, false, "sorry, endpoint is not found", "")
			return c.Status(fiber.StatusNotFound).JSON(webResponse)
		},
	)
}
