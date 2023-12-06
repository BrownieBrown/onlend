package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewUserRepository(db DBTX) models.UserRepository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	logger := utils.GetLogger()
	id := uuid.New()
	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) returning id"

	err := r.db.QueryRowContext(ctx, query, user.Id, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		logger.Error("Error while creating user", zap.Error(err))
		return &models.User{}, err
	}

	user.Id = id
	return user, nil
}
