package pkg

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/habibiiberahim/gofiber-clean-architecture/configs"
// 	databases "github.com/habibiiberahim/gofiber-clean-architecture/databases/migrations"
// 	"github.com/sirupsen/logrus"
// 	"github.com/urfave/cli/v2"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func runMigrate(database *gorm.DB) {
// 	for _, entity := range databases.RegisterEntities() {
// 		err := database.Debug().AutoMigrate(entity.Entity)
// 		if err != nil {
// 			logrus.Fatal(err)
// 		}
// 	}
// 	logrus.Info("Database migrate successfully")
// }

// func runSeeder(database *gorm.DB) {
// 	logrus.Info("Database seeder successfully")
// }

// func InitCommands(db *gorm.DB) {
// 	app := &cli.App{
// 		Commands: []*cli.Command{
// 			{
// 				Name:  "db:migrate",
// 				Usage: "migration database",
// 				Action: func(cCtx *cli.Context) error {
// 					runMigrate(db)
// 					return nil
// 				},
// 			},
// 			{
// 				Name:  "db:seed",
// 				Usage: "seeder database",
// 				Action: func(cCtx *cli.Context) error {
// 					runSeeder(db)
// 					return nil
// 				},
// 			},
// 		},
// 	}

// 	if err := app.Run(os.Args); err != nil {
// 		log.Fatal(err)
// 	}

// }
