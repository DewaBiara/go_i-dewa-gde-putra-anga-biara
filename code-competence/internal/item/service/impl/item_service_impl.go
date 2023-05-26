package impl

import (
	"context"

	"github.com/DewaBiara/Code-Competence/internal/item/dto"
	"github.com/DewaBiara/Code-Competence/internal/item/repository"
	"github.com/DewaBiara/Code-Competence/internal/item/service"
	"github.com/google/uuid"
)

type (
	ItemServiceImpl struct {
		itemRepository repository.ItemRepository
	}
)

func NewItemServiceImpl(itemRepository repository.ItemRepository) service.ItemService {
	return &ItemServiceImpl{
		itemRepository: itemRepository,
	}
}

func (u *ItemServiceImpl) CreateItem(ctx context.Context, item *dto.CreateItemRequest) error {

	itemEntity := item.ToEntity()
	itemEntity.ID = uuid.New().String()

	err := u.itemRepository.CreateItem(ctx, itemEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *ItemServiceImpl) GetSingleItem(ctx context.Context, itemID string) (*dto.GetSingleItemResponse, error) {
	item, err := d.itemRepository.GetSingleItem(ctx, itemID)
	if err != nil {
		return nil, err
	}

	var itemResponse = dto.NewGetSingleItemResponse(item)

	return itemResponse, nil
}

func (u *ItemServiceImpl) GetCategoryItem(ctx context.Context, categoryID, page int, limit int) (*dto.GetPageItemsResponse, error) {
	offset := (page - 1) * limit

	items, err := u.itemRepository.GetCategoryItem(ctx, categoryID, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageItemsResponse(items), nil
}

func (u *ItemServiceImpl) GetNameItem(ctx context.Context, name string, page int, limit int) (*dto.GetPageItemsResponse, error) {
	offset := (page - 1) * limit

	items, err := u.itemRepository.GetNameItem(ctx, name, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageItemsResponse(items), nil
}

func (u *ItemServiceImpl) GetPageItem(ctx context.Context, page int, limit int) (*dto.GetPageItemsResponse, error) {
	offset := (page - 1) * limit

	items, err := u.itemRepository.GetPageItem(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageItemsResponse(items), nil
}

func (u *ItemServiceImpl) UpdateItem(ctx context.Context, itemID string, updateItem *dto.UpdateItemRequest) error {
	item := updateItem.ToEntity()
	item.ID = itemID

	return u.itemRepository.UpdateItem(ctx, item)
}

func (d *ItemServiceImpl) DeleteItem(ctx context.Context, itemID string) error {

	return d.itemRepository.DeleteItem(ctx, itemID)
}
