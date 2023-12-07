package postgres

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func InitDB(cfg models.PostgresConfig) (*Database, error) {

	l, err := utils.NewZapLogger()
	if err != nil {
		return nil, err
	}
	logger := l.GetLogger()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	driver := "postgres"

	db, err := sql.Open(driver, dsn)
	if err != nil {
		logger.Error("Error connecting to postgres database", zap.Error(err))
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Error("Error on connecting to postgres database", zap.Error(err))
		return nil, err
	}

	return &Database{db: db}, nil
}

func (db *Database) Close() {
	db.Close()
}

func (db *Database) GetDB() *sql.DB {
	return db.db
}
