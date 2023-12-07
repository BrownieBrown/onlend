package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"server/internal/database/postgres"
	repo "server/internal/onlend/repo/postgres"
	"server/internal/onlend/rest"
	"server/internal/onlend/router"
	"server/internal/onlend/service"
	"server/internal/utils"
	"server/pkg/models"
)

func main() {
	utils.InitLogger()
	logger := utils.GetLogger()

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading.env file", zap.Error(err))
	}

	config := models.PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_NAME"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	}

	db, err := postgres.InitDB(config)
	if err != nil {
		logger.Error("Could not initialize db connection", zap.Error(err))
	}

	userRepository := repo.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepository)
	userHandler := rest.NewUserHandler(userService)
	r := router.NewRouter()

	r.InitRouter(userHandler)
	err = r.Start("0.0.0.0:8080")
	if err != nil {
		logger.Error("Could not start server", zap.Error(err))
		return
	}
}
