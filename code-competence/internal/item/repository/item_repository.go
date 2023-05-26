package repository

import (
	"context"

	"github.com/DewaBiara/Code-Competence/pkg/entity"
)

type ItemRepository interface {
	CreateItem(ctx context.Context, item *entity.Item) error
	UpdateItem(ctx context.Context, item *entity.Item) error
	GetSingleItem(ctx context.Context, itemID string) (*entity.Item, error)
	GetCategoryItem(ctx context.Context, categoryID, limit int, offset int) (*entity.Items, error)
	GetNameItem(ctx context.Context, name string, limit int, offset int) (*entity.Items, error)
	GetPageItem(ctx context.Context, limit int, offset int) (*entity.Items, error)
	DeleteItem(ctx context.Context, itemID string) error
}
