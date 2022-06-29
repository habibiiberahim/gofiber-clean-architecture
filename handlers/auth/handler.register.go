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
		validation.Field(&input.Fullname, validation.Required,),
		validation.Field(&input.Email, validation.Required),
	)
	
	if e != nil {
		jsonRes := helpers.APIResponse(ctx, "Invalid Input Data", fiber.StatusBadRequest, fiber.MethodPost, e)
		return ctx.Status(fiber.StatusBadRequest).JSON(jsonRes)
	}

	res, err := h.service.RegisterService(&input)
	switch err.Type{
	case "error_01":
		jsonRes := helpers.APIResponse(ctx, "Email already exist", fiber.StatusConflict, fiber.MethodPost, res)
		return ctx.Status(fiber.StatusAccepted).JSON(jsonRes)
	case "error_02":
		jsonRes := helpers.APIResponse(ctx, "Generate accessToken failed", fiber.StatusOK, fiber.MethodPost, res)
		return ctx.Status(fiber.StatusAccepted).JSON(jsonRes)
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, pkg.GodotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			jsonRes := helpers.APIResponse(ctx, "Generate accessToken failed", fiber.StatusBadRequest, fiber.MethodPost, nil)
			return ctx.Status(fiber.StatusAccepted).JSON(jsonRes)
		}

		// _, errSendMail := pkg.SendGridMail(res.Fullname, res.Email, "Activation Account", "template_register", accessToken)

		// if errSendMail != nil {
		// 	defer logrus.Error(errSendMail.Error())
		// 	helpers.APIResponse(ctx, "Sending email activation failed", fiber.StatusBadRequest, fiber.MethodPost, nil)
		// 	return
		// }
		
		jsonRes := helpers.APIResponse(ctx, "Register new account successfully", fiber.StatusCreated, fiber.MethodPost, accessToken)
		return ctx.Status(fiber.StatusCreated).JSON(jsonRes)
	}
}