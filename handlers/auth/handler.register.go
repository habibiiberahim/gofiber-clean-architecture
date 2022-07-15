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

func (h *handlerRegister) RegisterHandler(c *fiber.Ctx) error {
	//parsing from json to schema
	var input schemas.SchemaAuth
	if err := c.BodyParser(&input); err != nil {
		logrus.Debug(err.Error())
	}
	//validation again
	e := validation.ValidateStruct(&input,
		validation.Field(&input.Password, validation.Required),
		validation.Field(&input.Fullname, validation.Required),
		validation.Field(&input.Email, validation.Required),
	)

	if e != nil {
		helpers.APIResponse(c, fiber.StatusBadRequest, false, "Invalid Input Data", e)

	}

	res, err := h.service.RegisterService(&input)
	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, fiber.StatusConflict, false, "Email already exist", res)
		return nil
	case "error_02":
		helpers.APIResponse(c, fiber.StatusOK, false, "Generate accessToken failed", res)
		return nil
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, pkg.GodotEnv("JWT_SECRET_KEY"), 60)
		if errToken != nil {
			defer logrus.Error(errToken.Error())
			helpers.APIResponse(c, fiber.StatusBadRequest, false, "Generate accessToken failed", nil)
			return nil
		}
		// _, errSendMail := pkg.SendGridMail(res.Fullname, res.Email, "Activation Account", "template_register", accessToken)
		// if errSendMail != nil {
		// 	defer logrus.Error(errSendMail.Error())
		// 	helpers.APIResponse( "Sending email activation failed", fiber.StatusBadRequest,  nil)
		// 	return
		// }

		helpers.APIResponse(c, fiber.StatusCreated, true, "Register new account successfully", accessToken)
		return nil
	}
}
