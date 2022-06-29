package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
	services "github.com/habibiiberahim/gofiber-clean-architecture/services/auth"
)

type handlerPing struct {
	service services.ServicePing
}

func NewHandlerPing(service services.ServicePing) *handlerPing {
	return &handlerPing{service: service}
}

func (h *handlerPing) PingHandler (ctx *fiber.Ctx)  error {
	res := h.service.PingService()
	jsonRes := helpers.APIResponse(ctx, "Success", fiber.StatusOK, fiber.MethodGet, res)
	return ctx.Status(fiber.StatusAccepted).JSON(jsonRes)
}
