package services

import (
	"api/connection"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Addproduct(c *gin.Context) {
	type Product struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required"`
		Stock int     `json:"stock" binding:"required"`
	}

	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := connection.Mysql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()
	query := "INSERT INTO product (name, price, stock) VALUES (?, ?, ?)"
	result, err := db.Exec(query, newProduct.Name, newProduct.Price, newProduct.Stock)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to insert product",
			"details": err.Error(),
		})
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve last insert ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product added successfully!",
		"product": newProduct,
		"id":      lastInsertID,
	})
}
