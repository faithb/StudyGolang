package biz

import (
	"context"
	"social-todo-list/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(context context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(context context.Context, condition map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type UpdateItemBiz struct {
	storage UpdateItemStorage
}

func NewUpdateItemBiz(storage UpdateItemStorage) *UpdateItemBiz {

	return &UpdateItemBiz{storage: storage}
}

func (biz *UpdateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate model.TodoItemUpdate) error {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"id": id}, &dataUpdate); err != nil {
		return err
	}

	return nil
}
