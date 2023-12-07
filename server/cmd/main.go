package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"server/internal/database/postgres"
	repo "server/internal/onlend/repo/postgres"
	"server/internal/onlend/rest"
	"server/internal/onlend/router"
	"server/internal/onlend/service"
	"server/internal/utils"
	"server/pkg/models"
	"time"
)

func main() {
	l, err := utils.NewZapLogger()
	if err != nil {
		log.Fatal("Failed to initialize logger", err)
	}

	logger := l.GetLogger()

	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading.env file", zap.Error(err))
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
		logger.Fatal("Could not initialize db connection", zap.Error(err))
	}
	timeDuration := time.Duration(2) * time.Second

	userRepository := repo.NewUserRepository(db.GetDB(), l)
	userService := service.NewUserService(userRepository, l, timeDuration)
	userHandler := rest.NewUserHandler(userService, l)
	r := router.NewRouter()

	r.InitRouter(userHandler)
	err = r.Start(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		logger.Fatal("Could not start server", zap.Error(err))
	}
}
