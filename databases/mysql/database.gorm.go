package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/habibiiberahim/gofiber-clean-architecture/configs"
	databases "github.com/habibiiberahim/gofiber-clean-architecture/databases/migrations"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect sets the db client of database using configuration
func GetDatabase() (*gorm.DB, error) {
	cfg := configs.DBCfg()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	sqlDB, _ := sql.Open("mysql", dsn)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime * time.Minute))

	database, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		logrus.Panic(err)
	}
	return database, nil
}

func runMigrate(database *gorm.DB) {
	for _, entity := range databases.RegisterEntities() {
		err := database.Debug().AutoMigrate(entity.Entity)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	logrus.Info("Database migrate successfully")
}

func runSeeder(database *gorm.DB) {
	logrus.Info("Database seeder successfully")
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
		log.Fatal(err)
	}

}
