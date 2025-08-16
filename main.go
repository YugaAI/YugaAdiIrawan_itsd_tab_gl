package main

import (
	"interview/BackendAPIHumanCapital/internal/configs"
	"interview/BackendAPIHumanCapital/internal/handlers"
	"interview/BackendAPIHumanCapital/internal/models"
	"interview/BackendAPIHumanCapital/internal/repository"
	"interview/BackendAPIHumanCapital/internal/services"
	internalsql "interview/BackendAPIHumanCapital/pkg/internalSql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
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
