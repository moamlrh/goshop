package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/moamlrh/goshop/internal/models"
	"github.com/moamlrh/goshop/pkg/dtos"
)

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Add(ctx context.Context, user *models.User) (*models.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *"
	err := r.DB.QueryRowx(query, user.Username, user.Email, user.Password).StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE id=$1"
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM users"
	err := r.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetQueryable(ctx context.Context, q dtos.Queryable) ([]models.User, error) {
	var users []models.User
	query, args := q.BuildQueryWithFilters("SELECT * FROM users")
	err := r.DB.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4"
	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, user *models.User) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.DB.Exec(query, user.ID)
	if err != nil {
		return err
	}
	return nil
}
