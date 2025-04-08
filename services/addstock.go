package services

import (
	"api/connection"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Addstock(c *gin.Context) {
	type Product struct {
		Id    int     `json:"id" binding:"required"`
		Stock int     `json:"stock"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
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

	query := "UPDATE product SET stock = stock + ? WHERE id = ?"
	oldstock := "SELECT id, stock, name, price FROM product WHERE id = ?"

	// Ürünü güncelle
	result, err := db.Exec(query, newProduct.Stock, newProduct.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve affected rows"})
		return
	}

	rows, err := db.Query(oldstock, newProduct.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product information"})
		return
	}
	defer rows.Close()

	var updatedProduct Product
	if rows.Next() {
		err = rows.Scan(&updatedProduct.Id, &updatedProduct.Stock, &updatedProduct.Name, &updatedProduct.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan product information"})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// JSON yanıtında daha fazla bilgi döndür
	c.JSON(http.StatusOK, gin.H{
		"message":       "Product stock updated successfully!",
		"product":       updatedProduct,
		"rows_affected": rowsAffected,
	})
}
