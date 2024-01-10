package postgres

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"
)

func NewUserRepository(db DBTX, logger utils.Logger) models.UserRepository {
	return &repository{db: db, logger: logger}
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	logger := r.logger.GetLogger()

	newUUID := uuid.New()

	user.Id = newUUID

	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) returning id"
	var returnedId uuid.UUID
	err := r.db.QueryRowContext(ctx, query, user.Id, user.Username, user.Email, user.Password).Scan(&returnedId)
	if err != nil {
		logger.Error("Error while creating user", zap.Error(err))
		return nil, err
	}

	user.Id = returnedId
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	logger := r.logger.GetLogger()

	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		logger.Error("Error while finding user by email", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	logger := r.logger.GetLogger()

	query := "SELECT id, username, email, password FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		logger.Error("Error while getting all users", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			logger.Error("Error while scanning user", zap.Error(err))
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
