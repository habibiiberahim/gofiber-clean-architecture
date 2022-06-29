package handlers

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
	services "github.com/habibiiberahim/gofiber-clean-architecture/services/auth"
)

type handlerLogin struct {
	service services.ServiceLogin
}

func NewHandlerLogin(service services.ServiceLogin)*handlerLogin {
	return &handlerLogin{
		service: service,
	}
}

func (h *handlerLogin)LoginHandler(ctx *fiber.Ctx) error {
	var input schemas.SchemaAuth
	if err := ctx.BodyParser(&input); err != nil {
		logrus.Debug(err.Error())
	}

	
	e :=validation.ValidateStruct(&input,
		validation.Field(&input.Email,validation.Required),
		validation.Field(&input.Password,validation.Required),
	)
	fmt.Println(e)
	if e != nil {
		jsonRes := helpers.APIResponse(ctx, "Invalid Input Data", fiber.StatusBadRequest, fiber.MethodPost, e)
		return ctx.Status(fiber.StatusBadRequest).JSON(jsonRes)
	}

	res, err := h.service.LoginService(&input)

	switch err.Type{
	case "error_01" : 
		jsonRes := helpers.APIResponse(ctx, "User account is not registered", err.Code, fiber.MethodPost, res)
		return ctx.Status(err.Code).JSON(jsonRes)
	case "error_02" : 
		jsonRes := helpers.APIResponse(ctx, "User account is not active", err.Code, fiber.MethodPost, res)
		return ctx.Status(err.Code).JSON(jsonRes)
	case "error_03" : 
		jsonRes := helpers.APIResponse(ctx, "Username or password is wrong", err.Code, fiber.MethodPost, res)
		return ctx.Status(err.Code).JSON(jsonRes)
	default : 
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, pkg.GodotEnv("JWT_SECRET"), 24*60*1)

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
		
		jsonRes := helpers.APIResponse(ctx, "Login successfully", fiber.StatusOK, fiber.MethodPost, accessToken)
		return ctx.Status(fiber.StatusCreated).JSON(jsonRes)
	}
}