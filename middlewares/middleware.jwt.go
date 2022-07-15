package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("Authorization") == "" {
			c.Status(fiber.StatusBadRequest)
			helpers.APIResponse(c, fiber.StatusBadRequest, false, "Missing or malformed JWT", "")
		}
		token, err := pkg.VerifyTokenHeader(c, pkg.GodotEnv("JWT_SECRET_KEY"))
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			helpers.APIResponse(c, fiber.StatusUnauthorized, false, "Invalid or expired JWT", "")
		}
		cookie := new(fiber.Cookie)
		cookie.Name = "Jwt"
		cookie.Value = token.Raw
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.Cookie(cookie)
		return c.Status(fiber.StatusOK).Next()
	}
}
