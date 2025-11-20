package main

import (
	"log"
	"net/http"
	"os"
	ginitem "social-todo-list/modules/item/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_CONNECT")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// CRUD : Create, Read, Update, Delete
	// POST /v1/items (Create a new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT || PATCH) /v1/items/:id (update item by id)
	// DELETE /v1/items/:id (delete item by id)

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItems(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PUT("/:id", ginitem.UpdateItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	} // Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "item",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":3001"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
