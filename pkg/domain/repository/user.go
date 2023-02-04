package repository

import (
	"context"

	"github.com/FarStep131/go-jwt/pkg/domain/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}
