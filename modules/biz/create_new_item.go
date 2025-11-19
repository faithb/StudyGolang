package biz

import (
	"context"
	"social-todo-list/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(context context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
	storage CreateItemStorage
}

func NewCreateItemBiz(storage CreateItemStorage) *createItemBiz {
	return &createItemBiz{storage: storage}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
