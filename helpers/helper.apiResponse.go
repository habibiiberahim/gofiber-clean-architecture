package helpers

import (
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
)

func APIResponse(Code int, Success bool, Message string, Data interface{}) interface{} {

	jsonResponse := schemas.SchemaResponses{
		Code:    Code,
		Success: Success,
		Message: Message,
		Data:    Data,
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
