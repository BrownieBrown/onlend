package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"server/pkg/models"
)

func NewUserRepository(db *sql.DB) models.UserRepository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.Id = uuid.New()

	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) returning id"
	var returnedId uuid.UUID
	err := r.db.QueryRowContext(ctx, query, user.Id, user.Username, user.Email, user.Password).Scan(&returnedId)
	if err != nil {
		return nil, err
	}

	user.Id = returnedId
	return user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	query := "SELECT id, username, email, password FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
