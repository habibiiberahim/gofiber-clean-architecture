package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/habibiiberahim/gofiber-clean-architecture/handlers/auth"
	repositorys "github.com/habibiiberahim/gofiber-clean-architecture/repositorys/auth"
	services "github.com/habibiiberahim/gofiber-clean-architecture/services/auth"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *fiber.App) {
	
	/**
	@description All Handler Auth
	*/
	pingRepository := repositorys.NewRepositoryPing(db)
	pingService := services.NewServicePing(pingRepository)
	pingHandler := handlers.NewHandlerPing(pingService)

	registerRepository := repositorys.NewRepositoryRegister(db)
	registerService := services.NewServiceRegister(registerRepository)
	registerHandler := handlers.NewHandlerRegister(registerService)

	loginRepository := repositorys.NewRepositoryLogin(db)
	loginService := services.NewServiceLogin(loginRepository)
	loginHandler := handlers.NewHandlerLogin(loginService)
	
	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("api/v1")
	groupRoute.Get("/ping", pingHandler.PingHandler)
	groupRoute.Post("/register", registerHandler.RegisterHandler)
	groupRoute.Post("/login", loginHandler.LoginHandler)
	// groupRoute.Post("/activation/:token", activationHandler.ActivationHandler)
	// groupRoute.Post("/resend-token", resendHandler.ResendHandler)
	// groupRoute.Post("/forgot-password", forgotHandler.ForgotHandler)
	// groupRoute.Post("/change-password/:token", resetHandler.ResetHandler)pr
}