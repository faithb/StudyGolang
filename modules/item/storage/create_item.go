package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (sql *SqlStorage) CreateItem(context context.Context, data *model.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
