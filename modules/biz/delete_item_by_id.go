package biz

import (
	"context"
	"social-todo-list/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(context context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(context context.Context, condition map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type DeleteItemBiz struct {
	storage DeleteItemStorage
}

func NewDeleteItemBiz(storage DeleteItemStorage) *DeleteItemBiz {

	return &DeleteItemBiz{storage: storage}
}

func (biz *DeleteItemBiz) DeleteItemById(ctx context.Context, id int) error {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	deletedStatus := model.ItemStatusDeleted
	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"id": id}, &model.TodoItemUpdate{
		Status: (&deletedStatus).String(),
	}); err != nil {
		return err
	}

	return nil
}
