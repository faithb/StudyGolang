package model

import (
	"errors"
	"social-todo-list/common"
)

var (
	ErrTitleIsBlank = errors.New("title can't be blank")
	ErrItemDeleted  = errors.New("item deleted and cannot be modified")
)

type TodoItem struct {
	common.SqlModel
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id;primaryKey;autoIncrement;"`
	Title       string      `json:"title" binding:"required" gorm:"column:title;"`
	Description string      `json:"description" binding:"required" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       string  `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      string  `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
