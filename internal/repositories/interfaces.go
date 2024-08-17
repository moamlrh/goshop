package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/moamlrh/goshop/internal/models"
	"github.com/moamlrh/goshop/pkg/dtos"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetQueryable(ctx context.Context, queryable dtos.Queryable) ([]models.User, error)
	Add(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	Delete(ctx context.Context, user *models.User) error
}
