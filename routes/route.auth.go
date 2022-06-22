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

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("api/v1")
	groupRoute.Get("/ping", pingHandler.PingHandler)
	// groupRoute.POST("/register", registerHandler.RegisterHandler)
	// groupRoute.POST("/login", loginHandler.LoginHandler)
	// groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)
	// groupRoute.POST("/resend-token", resendHandler.ResendHandler)
	// groupRoute.POST("/forgot-password", forgotHandler.ForgotHandler)
	// groupRoute.POST("/change-password/:token", resetHandler.ResetHandler)pr
}