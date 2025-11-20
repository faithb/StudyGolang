package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (sql *SqlStorage) UpdateItem(context context.Context, condition map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := sql.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
