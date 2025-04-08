package services

import (
	"api/addons"
	_ "api/connection"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func Login(c *gin.Context) {
	type login struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var user login
	input := user.Password

	hash := addons.Encrypt(input)
	fmt.Printf("Your secure hash: %s\n", hash)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}
