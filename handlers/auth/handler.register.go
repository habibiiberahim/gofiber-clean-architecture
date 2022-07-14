package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
	services "github.com/habibiiberahim/gofiber-clean-architecture/services/auth"
	"github.com/sirupsen/logrus"
)

type handlerRegister struct {
	service services.ServiceRegister
}

func NewHandlerRegister(service services.ServiceRegister) *handlerRegister {
	return &handlerRegister{
		service: service,
	}
}

func (h *handlerRegister) RegisterHandler(ctx *fiber.Ctx) error {
	//parsing from json to schema
	var input schemas.SchemaAuth
	if err := ctx.BodyParser(&input); err != nil {
		logrus.Debug(err.Error())
	}
	//validation again
	e := validation.ValidateStruct(&input,
		validation.Field(&input.Password, validation.Required),
		validation.Field(&input.Fullname, validation.Required),
		validation.Field(&input.Email, validation.Required),
	)

	if e != nil {
		webResponse := helpers.APIResponse(fiber.StatusBadRequest, false, "Invalid Input Data", e)
		return ctx.Status(fiber.StatusBadRequest).JSON(webResponse)
	}

	res, err := h.service.RegisterService(&input)
	switch err.Type {
	case "error_01":
		webResponse := helpers.APIResponse(fiber.StatusConflict, false, "Email already exist", res)
		return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
	case "error_02":
		webResponse := helpers.APIResponse(fiber.StatusOK, false, "Generate accessToken failed", res)
		return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, pkg.GodotEnv("JWT_SECRET_KEY"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			webResponse := helpers.APIResponse(fiber.StatusBadRequest, false, "Generate accessToken failed", nil)
			return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
		}

		// _, errSendMail := pkg.SendGridMail(res.Fullname, res.Email, "Activation Account", "template_register", accessToken)

		// if errSendMail != nil {
		// 	defer logrus.Error(errSendMail.Error())
		// 	helpers.APIResponse( "Sending email activation failed", fiber.StatusBadRequest,  nil)
		// 	return
		// }

		webResponse := helpers.APIResponse(fiber.StatusCreated, true, "Register new account successfully", accessToken)
		return ctx.Status(fiber.StatusCreated).JSON(webResponse)
	}
}
