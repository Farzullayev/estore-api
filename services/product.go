package services

import (
	"database/sql"
	"net/http"
	"strconv"

	"api/connection"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Product(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	db, err := connection.Mysql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	query := "SELECT id, name, price, stock FROM product WHERE id = ?"
	row := db.QueryRow(query, id)

	type Product struct {
		Id    int     `json:"id"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}

	var product Product
	err = row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product found successfully!",
		"product": product,
	})
}
