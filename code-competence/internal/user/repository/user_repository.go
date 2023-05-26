package repository

import (
	"context"

	"github.com/DewaBiara/Code-Competence/pkg/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	GetBriefUsers(ctx context.Context, limit int, offset int) (*entity.Users, error)
	UpdateUser(ctx context.Context, user *entity.User) error
}
