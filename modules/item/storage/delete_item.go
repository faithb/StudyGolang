package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (sql *SqlStorage) DeleteItem(context context.Context, condition map[string]interface{}, dataDelete *model.TodoItemUpdate) error {
	if err := sql.db.Where(condition).Updates(dataDelete).Error; err != nil {
		return err
	}

	return nil
}
