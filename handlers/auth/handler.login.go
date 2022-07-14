package handlers

import (
	"fmt"
	"time"

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

func NewHandlerLogin(service services.ServiceLogin) *handlerLogin {
	return &handlerLogin{
		service: service,
	}
}

func (h *handlerLogin) LoginHandler(c *fiber.Ctx) error {
	var input schemas.SchemaAuth
	if err := c.BodyParser(&input); err != nil {
		logrus.Debug(err.Error())
	}

	e := validation.ValidateStruct(&input,
		validation.Field(&input.Email, validation.Required),
		validation.Field(&input.Password, validation.Required),
	)
	fmt.Println(e)
	if e != nil {
		jsonRes := helpers.APIResponse(fiber.StatusBadRequest, false, "Invalid Input Data", e)
		return c.Status(fiber.StatusBadRequest).JSON(jsonRes)
	}

	res, err := h.service.LoginService(&input)

	switch err.Type {
	case "error_01":
		webResponse := helpers.APIResponse(err.Code, false, "User account is not registered", res)
		return c.Status(err.Code).JSON(webResponse)
	case "error_02":
		webResponse := helpers.APIResponse(err.Code, false, "User account is not active", res)
		return c.Status(err.Code).JSON(webResponse)
	case "error_03":
		webResponse := helpers.APIResponse(err.Code, false, "Username or password is wrong", res)
		return c.Status(err.Code).JSON(webResponse)
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, pkg.GodotEnv("JWT_SECRET_KEY"), 24*60*1)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			webResponse := helpers.APIResponse(fiber.StatusBadRequest, true, "Generate accessToken failed", nil)
			return c.Status(fiber.StatusAccepted).JSON(webResponse)
		}

		// _, errSendMail := pkg.SendGridMail(res.Fullname, res.Email, "Activation Account", "template_register", accessToken)

		// if errSendMail != nil {
		// 	defer logrus.Error(errSendMail.Error())
		// 	helpers.APIResponse(c, "Sending email activation failed", fiber.StatusBadRequest, fiber.MethodPost, nil)
		// 	return
		// }

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    accessToken,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		webResponse := helpers.APIResponse(fiber.StatusOK, true, "Login successfully", accessToken)
		return c.Status(fiber.StatusCreated).JSON(webResponse)
	}
}
