package main

import (
    "fmt"
	"database/sql"
    _"github.com/go-sql-driver/mysql"
	"api/services"
	"github.com/gin-gonic/gin"
	_"api/connection"
)
func database() {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8163)/api")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    fmt.Println("Connected to Database!")

    // 'product' tablosunu kontrol edip oluşturma
    query := `
        CREATE TABLE IF NOT EXISTS product (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            price DECIMAL(10, 2) NOT NULL,
            stock INT NOT NULL DEFAULT 0, -- Stok bilgisini ekliyoruz
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
	r.Run()
}