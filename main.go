package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"main.go/intranal/configs"
	"main.go/intranal/handlers"
	"main.go/intranal/models"
	"main.go/intranal/repository"
	"main.go/intranal/services"
	"main.go/pkg/internalsql"
)

func main() {

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./intranal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("error initializing config: %v", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&models.Employee{})

	r := gin.Default()

	repo := repository.NewRepository(db)
	svc := services.NewService(repo)
	h := handlers.NewHandler(r, svc)
	h.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
