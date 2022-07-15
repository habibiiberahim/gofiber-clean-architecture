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

func (h *handlerPing) PingHandler(c *fiber.Ctx) error {
	res := h.service.PingService()
	helpers.APIResponse(c, fiber.StatusAccepted, true, "", res)
	return nil
}
