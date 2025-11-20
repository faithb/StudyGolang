package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		context context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type ListItemBiz struct {
	storage ListItemStorage
}

func NewListItemBiz(storage ListItemStorage) *ListItemBiz {
	return &ListItemBiz{storage: storage}
}

func (biz *ListItemBiz) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.TodoItem, error) {
	data, err := biz.storage.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
