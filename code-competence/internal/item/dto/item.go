package dto

import (
	"github.com/DewaBiara/Code-Competence/pkg/entity"
)

type CreateItemRequest struct {
	Name        string `json:"name" validate:"required"`
	CategoryID  int    `json:"categoryid" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdby"`
}

func (u *CreateItemRequest) ToEntity() *entity.Item {
	return &entity.Item{
		Name:        u.Name,
		CategoryID:  u.CategoryID,
		Price:       u.Price,
		Stock:       u.Stock,
		Description: u.Description,
		CreatedBy:   u.CreatedBy,
	}
}

type UpdateItemRequest struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	CategoryID  int    `json:"categoryid" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	UpdatedBy   string `json:"updatedby"`
}

func (u *UpdateItemRequest) ToEntity() *entity.Item {
	return &entity.Item{
		Name:        u.Name,
		CategoryID:  u.CategoryID,
		Price:       u.Price,
		Stock:       u.Stock,
		Description: u.Description,
		UpdatedBy:   u.UpdatedBy,
	}
}

type GetSingleItemResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	CategoryID  int    `json:"categoryid" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdby"`
	UpdatedBy   string `json:"updatedby"`
}

func NewGetSingleItemResponse(item *entity.Item) *GetSingleItemResponse {
	return &GetSingleItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		CategoryID:  item.CategoryID,
		Price:       item.Price,
		Stock:       item.Stock,
		Description: item.Description,
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
	}
}

type GetPageItemResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	CategoryID  int    `json:"categoryid" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdby"`
	UpdatedBy   string `json:"updatedby"`
}

func NewGetPageItemResponse(item *entity.Item) *GetPageItemResponse {
	return &GetPageItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		CategoryID:  item.CategoryID,
		Price:       item.Price,
		Stock:       item.Stock,
		Description: item.Description,
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
	}
}

type GetPageItemsResponse []GetPageItemResponse

func NewGetPageItemsResponse(items *entity.Items) *GetPageItemsResponse {
	var getPageItems GetPageItemsResponse
	for _, items := range *items {
		getPageItems = append(getPageItems, *NewGetPageItemResponse(&items))
	}
	return &getPageItems
}
