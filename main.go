package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/configs"
	databases "github.com/habibiiberahim/gofiber-clean-architecture/databases/mysql"
	"github.com/habibiiberahim/gofiber-clean-architecture/middlewares"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"github.com/habibiiberahim/gofiber-clean-architecture/routes"
)

func main() {
	// Load all configs
	configs.LoadAllConfigs(".env")
	appCfg := configs.AppCfg()

	pkg.SetupLogger()
	logr := pkg.GetLogger()

	// connect to DB
	db, err := databases.GetDatabase()
	if err != nil {
		logr.Panicf("failed database setup. error: %v", err)
	}

	//Define Fibrt config & app
	fiberCfg := configs.FiberConfig()
	app := fiber.New(fiberCfg)

	// Attach Middlewares.
	middlewares.FiberMiddleware(app)

	// Routes.
	routes.InitAuthRoutes(db, app)
	routes.NotFoundRoute(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		logr.Infoln("Shutting down server...")
		_ = app.Shutdown()

	}()

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}
}
