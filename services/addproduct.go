package services

import (
	"api/connection"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Addproduct(c *gin.Context) {
	type Product struct {
		Name     string  `json:"name" binding:"required"`
		Category string  `json:"category" binding:"required"`
		Price    float64 `json:"price" binding:"required"`
		Stock    int     `json:"stock" binding:"required"`
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

	var categoryName string
	err = db.QueryRow("SELECT name FROM category WHERE name = ?", newProduct.Category).Scan(&categoryName)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category does not exist"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check category"})
		return
	}

	query := "INSERT INTO product (name, category, price, stock) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, newProduct.Name, newProduct.Category, newProduct.Price, newProduct.Stock)

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
