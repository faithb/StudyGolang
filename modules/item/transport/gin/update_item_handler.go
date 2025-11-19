package ginitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		business := biz.NewUpdateItemBiz(storage.NewSQLStorage(db))
		if err := business.UpdateItemById(c.Request.Context(), id, data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
