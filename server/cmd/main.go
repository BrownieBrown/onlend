package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"server/internal/database/postgres"
	repo "server/internal/onlend/repo/postgres"
	"server/internal/onlend/rest"
	"server/internal/onlend/router"
	"server/internal/onlend/service"
	"server/internal/utils"
)

func main() {
	l, err := utils.NewZapLogger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		return
	}

	logger := l.GetLogger()

	err = godotenv.Load()
	if err != nil {
		logger.Fatal("Could not load env file", zap.Error(err))
	}

	config, err := utils.LoadConfig()
	if err != nil {
		logger.Fatal("Could not load postgres config", zap.Error(err))
	}

	db, err := postgres.InitDB(config.Postgres)

	if err != nil {
		logger.Fatal("Could not initialize db connection", zap.Error(err))
	}

	// repository init
	userRepository := repo.NewUserRepository(db.GetDB())
	accountRepository := repo.NewAccountRepository(db.GetDB())
	transactionRepository := repo.NewTransactionRepository(db.GetDB())

	// service init
	accountService := service.NewAccountService(accountRepository, config)
	userService := service.NewUserService(userRepository, accountService, config)
	transactionService := service.NewTransactionService(transactionRepository, accountService, config)

	// handler init
	userHandler := rest.NewUserHandler(userService, l, config)
	accountHandler := rest.NewAccountHandler(accountService, l, config)
	transactionHandler := rest.NewTransactionHandler(transactionService, l, config)

	// router init
	r := router.NewRouter()
	r.InitRouter(userHandler, accountHandler, transactionHandler)

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "localhost:8081"
	}
	err = r.Start(serverAddress)
	if err != nil {
		logger.Fatal("Could not start server", zap.Error(err))
		return
	}

	defer func(db *postgres.PSQLDatabase) {
		err := db.Close()
		if err != nil {
			logger.Fatal("Could not close db connection", zap.Error(err))
		}
	}(db)
}
