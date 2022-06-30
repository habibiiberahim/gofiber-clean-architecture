package databases

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/habibiiberahim/gofiber-clean-architecture/configs"
	databases "github.com/habibiiberahim/gofiber-clean-architecture/databases/migrations"
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
		return nil, fmt.Errorf("can't connected to database, %w", err)
	}

	for _, entity := range databases.RegisterEntities() {
		err := database.Debug().AutoMigrate(entity.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return database, nil
}

func InitCommands(db *gorm.DB) {
	cmdApp := cli.NewApp()
	cmdApp.Commands = []*cli.Command{
		{
			Name: "db:migrate",
			Action: func(ctx *cli.Context) error {
				fmt.Println("db migrate")
				return nil
			},
		},
	}
}
