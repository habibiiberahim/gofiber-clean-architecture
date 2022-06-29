package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/habibiiberahim/gofiber-clean-architecture/helpers"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"github.com/habibiiberahim/gofiber-clean-architecture/routes"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger := logrus.New()
	logger.Println("Starting App")
	app := SetupRouter()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c
		logrus.Println("Graefully shutting down")
		_ = app.Shutdown()

	}()

	if err := app.Listen(":" + pkg.GodotEnv("GO_PORT")); err != nil {
		logrus.Panic(err)
	}

	logrus.Println("Running cleanup task")
}

func SetupRouter() *fiber.App {
	db, err := SetupDatabase()
	if err != nil {
		logrus.Println(err)
	}
	app := fiber.New()
	// Use global middlewares.
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(ctx *fiber.Ctx) error {
			jsonRes := helpers.APIResponse(ctx, "You have requested too many", fiber.StatusTooManyRequests, fiber.MethodGet, "")
			return ctx.Status(fiber.StatusAccepted).JSON(jsonRes)
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	//routes init db and app
	routes.InitAuthRoutes(db, app)

	app.Use(func(ctx *fiber.Ctx) error {
		jsonRes := helpers.APIResponse(ctx, "This endpoint not found", fiber.StatusNotFound, fiber.MethodGet, "")
		return ctx.Status(fiber.StatusAccepted).JSON(jsonRes)
	})

	return app
}

func SetupDatabase() (*gorm.DB, error) {
	//create connection to database
	user := pkg.GodotEnv("DATABASE_USER")
	pass := pkg.GodotEnv("DATABASE_PASS")
	host := pkg.GodotEnv("DATABASE_HOST")
	port := pkg.GodotEnv("DATABASE_PORT")
	dbname := pkg.GodotEnv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	sqlDB, _ := sql.Open("mysql", dsn)
	maxIdleConn, _ := strconv.Atoi(pkg.GodotEnv("DB_MAX_IDLE_CONNECTION"))
	maxOpenConn, _ := strconv.Atoi(pkg.GodotEnv("DB_MAX_OPEN_CONNECTION"))
	maxLifetimeConn, _ := strconv.Atoi(pkg.GodotEnv("DB_MAX_LIFETIME_CONNECTION"))

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn) * time.Minute)

	database, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return database, err
}
