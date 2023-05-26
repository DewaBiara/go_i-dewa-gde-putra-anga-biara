package service

import (
	"context"

	"github.com/DewaBiara/Code-Competence/internal/item/dto"
)

type ItemService interface {
	CreateItem(ctx context.Context, item *dto.CreateItemRequest) error
	UpdateItem(ctx context.Context, itemID string, updateItem *dto.UpdateItemRequest) error
	GetSingleItem(ctx context.Context, itemID string) (*dto.GetSingleItemResponse, error)
	GetCategoryItem(ctx context.Context, categoryID, limit int, offset int) (*dto.GetPageItemsResponse, error)
	GetNameItem(ctx context.Context, name string, limit int, offset int) (*dto.GetPageItemsResponse, error)
	GetPageItem(ctx context.Context, limit int, offset int) (*dto.GetPageItemsResponse, error)
	DeleteItem(ctx context.Context, itemID string) error
}
