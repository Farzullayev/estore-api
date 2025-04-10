package main

import (
	_ "api/connection"
	"api/services"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func database() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8163)/api")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Connected to Database!")
	query := `
        CREATE TABLE IF NOT EXISTS product (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            price DECIMAL(10, 2) NOT NULL,
            stock INT NOT NULL DEFAULT 0,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `

	_, err = db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Product table checked or created successfully!")
}

func main() {
	database()
	r := gin.Default()
	r.GET("/ping", services.Ping)
	r.POST("/addproduct", services.Addproduct)
	r.DELETE("/deleteproduct", services.Deleteproduct)
	r.GET("/getproduct", services.Getproduct)
	r.POST("/addstock", services.Addstock)
	r.POST("/login", services.Login)
	r.Run()

}
