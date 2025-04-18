package services

import (
	"api/connection"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Addcategory(c *gin.Context) {
	type category struct {
		Name string `json:"name" binding:"required"`
	}

	var newCategory category

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connection.Mysql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	var existingCategory string
	err = db.QueryRow("SELECT name FROM category WHERE name = ?", newCategory.Name).Scan(&existingCategory)

	if err == nil {
	
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category with this name already exists"})
		return
	}

	if err == sql.ErrNoRows {
	
		_, err := db.Exec("INSERT INTO category (name) VALUES (?)", newCategory.Name)
		if err != nil {
			log.Println("Failed to insert category:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert category"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Category added successfully!",
			"category": newCategory,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
}
