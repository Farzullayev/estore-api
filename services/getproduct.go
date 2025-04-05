package services

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"api/connection"
)

func Getproduct(c *gin.Context) {
	type Product struct {
		Id    int     `json:"id" binding:"required"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}

	var getProduct Product

	if err := c.ShouldBindJSON(&getProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connection.Mysql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	query := "SELECT id, name, price, stock FROM product WHERE id = ?"
	row := db.QueryRow(query, getProduct.Id)

	err = row.Scan(&getProduct.Id, &getProduct.Name, &getProduct.Price, &getProduct.Stock)
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
		"product": getProduct,
	})
}