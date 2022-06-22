package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
)

func APIResponse(ctx *fiber.Ctx, Message string, StatusCode int, Method string, Data interface{}) interface{}{

	jsonResponse := schemas.SchemaResponses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	return jsonResponse
}

// func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
// 	errResponse := schemas.SchemaErrorResponse{
// 		StatusCode: StatusCode,
// 		Method:     Method,
// 		Error:      Error,
// 	}

// 	ctx.AbortWithStatusJSON(StatusCode, errResponse)
// }