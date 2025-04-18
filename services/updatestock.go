package services

import (
	"api/connection"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Addstock(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	type Product struct {
		ID    int     `json:"id"`
		Stock int     `json:"stock"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	var newStock struct {
		Stock int `json:"stock"`
	}

	if err := c.ShouldBindJSON(&newStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connection.Mysql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	updateQuery := "UPDATE product SET stock = stock + ? WHERE id = ?"
	result, err := db.Exec(updateQuery, newStock.Stock, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve affected rows"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	selectQuery := "SELECT id, stock, name, price FROM product WHERE id = ?"
	row := db.QueryRow(selectQuery, id)

	var updatedProduct Product
	err = row.Scan(&updatedProduct.ID, &updatedProduct.Stock, &updatedProduct.Name, &updatedProduct.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan updated product information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Product stock updated successfully!",
		"product":       updatedProduct,
		"rows_affected": rowsAffected,
	})
}
