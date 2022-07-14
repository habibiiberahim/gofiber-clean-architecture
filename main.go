package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/entities"
	"github.com/habibiiberahim/gofiber-clean-architecture/middlewares"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"github.com/habibiiberahim/gofiber-clean-architecture/routes"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	pkg.SetupLogger()
	logr := pkg.GetLogger()

	// connect to DB
	db := SetupDatabase()

	// Checking command for migrating
	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		InitCommands(db)
	}

	//Define Fiber config & app
	app := fiber.New()

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
	appHost := pkg.GodotEnv("APP_HOST")
	appPort, _ := strconv.Atoi(pkg.GodotEnv("APP_PORT"))

	serverAddr := fmt.Sprintf("%s:%d", appHost, appPort)
	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}

}

func SetupDatabase() *gorm.DB {
	user := pkg.GodotEnv("DB_USER")
	password := pkg.GodotEnv("DB_PASSWORD")
	host := pkg.GodotEnv("DB_HOST")
	port := pkg.GodotEnv("DB_PORT")
	name := pkg.GodotEnv("DB_NAME")
	maxIdleConn, _ := strconv.Atoi(pkg.GodotEnv("DB_MAX_IDLE_CONNECTIONS"))
	maxOpenConn, _ := strconv.Atoi(pkg.GodotEnv("DB_MAX_OPEN_CONNECTIONS"))
	timeOut, _ := strconv.Atoi(pkg.GodotEnv("DB_MAX_LIFETIME_CONNECTIONS"))
	maxConnLifetime := time.Duration(timeOut) * time.Second

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	sqlDB, _ := sql.Open("mysql", dsn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxConnLifetime * time.Minute))

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		logrus.Panicf("failed database setup. error: %v", err)
	}
	return db
}

func InitCommands(db *gorm.DB) {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "db:migrate",
				Usage: "migration database",
				Action: func(cCtx *cli.Context) error {
					runMigrate(db)
					return nil
				},
			},
			{
				Name:  "db:seed",
				Usage: "seeder database",
				Action: func(cCtx *cli.Context) error {
					runSeeder(db)
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func runMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.User{},
	)
	if err != nil {
		logrus.Fatal("Database migrate unsuccessfully %v", err)
	}

	logrus.Info("Database migrate successfully")
}

func runSeeder(db *gorm.DB) {
	logrus.Info("Database seeder successfully")
}
