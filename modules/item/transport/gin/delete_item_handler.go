package ginitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/biz"
	"social-todo-list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewDeleteItemBiz(store)

		if err := business.DeleteItemById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
