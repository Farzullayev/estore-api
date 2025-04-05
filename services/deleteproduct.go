package services

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"api/connection"
)

func Deleteproduct(c *gin.Context) {
	type Product struct {
		Id int `json:"id" binding:"required"`
	}

	var remProduct Product

	if err := c.ShouldBindJSON(&remProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connection.Mysql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	query := "DELETE FROM product WHERE id = ?"
	result, err := db.Exec(query, remProduct.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed successfully!", "product": remProduct})
}