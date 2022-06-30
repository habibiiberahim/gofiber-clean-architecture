package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-jwt/jwt/v4"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		cors.New(),
		compress.New(),
		etag.New(),
		favicon.New(),
		recover.New(),
		limiter.New(limiter.Config{
			Max: 100,
			LimitReached: func(ctx *fiber.Ctx) error {
				jsonRes := helpers.APIResponse(ctx, "You have requested too many", fiber.StatusTooManyRequests, fiber.MethodGet, "")
				return ctx.Status(fiber.StatusAccepted).JSON(jsonRes)
			},
		}),
	)
}

func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin, ok := claims["admin"]
	if !ok || !isAdmin.(bool) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"msg": "Forbidden",
		})
	}

	return c.Next()
}
