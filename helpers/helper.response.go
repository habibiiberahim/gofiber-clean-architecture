package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
)

func APIResponse(c *fiber.Ctx, Code int, Success bool, Message string, Data interface{}) {

	jsonResponse := schemas.SchemaResponses{
		Code:    Code,
		Success: Success,
		Message: Message,
		Data:    Data,
	}
	c.JSON(jsonResponse)
}
