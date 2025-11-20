package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (sql *SqlStorage) ListItem(
	context context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	var items []model.TodoItem

	db := sql.db.Where("status != ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.
		TableName()).
		Count(&paging.Total).
		Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&items).
		Error; err != nil {
		return nil, err
	}

	return items, nil
}
