package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
)

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			helpers.APIResponse(c, fiber.StatusNotFound, false, "sorry, endpoint is not found", "")
			return nil
		},
	)
}
