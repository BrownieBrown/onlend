package postgres

import (
	"database/sql"
	"fmt"
	"server/pkg/models"

	_ "github.com/lib/pq"
)

type Database interface {
	Close() error
	GetDB() *sql.DB
	Open(driverName, dataSourceName string) error
	Ping() error
}

type PSQLDatabase struct {
	db *sql.DB
}

func NewPSQLDatabase() *PSQLDatabase {
	return &PSQLDatabase{}
}

func (db *PSQLDatabase) Close() error {
	return db.db.Close()
}

func (db *PSQLDatabase) GetDB() *sql.DB {
	return db.db
}

func (db *PSQLDatabase) Sync() error {
	return nil
}

func (db *PSQLDatabase) Open(driverName, dataSourceName string) error {
	var err error
	db.db, err = sql.Open(driverName, dataSourceName)
	return err
}

func (db *PSQLDatabase) Ping() error {
	err := db.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func InitDB(cfg models.PostgresConfig) (*PSQLDatabase, error) {
	dsn := buildDSN(cfg)

	psqlDB := NewPSQLDatabase()
	err := psqlDB.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = psqlDB.Ping(); err != nil {
		return nil, err
	}

	return psqlDB, nil
}

func buildDSN(cfg models.PostgresConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
}