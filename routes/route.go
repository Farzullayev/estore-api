package routes

import (
	"api/services"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.GET("/product/:id", services.Product)
	r.POST("/product", services.Addproduct)
	r.DELETE("/product/:id", services.Deleteproduct)
	r.POST("/category", services.Addcategory)
	r.PUT("/product/:id", services.Addstock)
	r.Run()
}
