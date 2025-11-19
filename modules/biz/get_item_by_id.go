package biz

import (
	"context"
	"social-todo-list/modules/item/model"
)

type GetItemStorage interface {
	GetItem(context context.Context, condition map[string]interface{}) (*model.TodoItem, error)
}

type getItemBiz struct {
	storage GetItemStorage
}

func NewGetItemBiz(storage GetItemStorage) *getItemBiz {
	return &getItemBiz{storage: storage}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
