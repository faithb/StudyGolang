package ginitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		business := biz.NewCreateItemBiz(storage.NewSQLStorage(db))

		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
