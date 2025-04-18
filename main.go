package main

import (
	_ "api/connection"
	"api/routes"
	"database/sql"
	"fmt"

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
			category VARCHAR(255) NOT NULL,
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

	routes.Routes()
	//r.GET("/ping", services.Ping)
	//r.POST("/addproduct", services.Addproduct)
	////r.GET("/product/:id", services.Product)
	//r.DELETE("/deleteproduct", services.Deleteproduct)
	///r.POST("/addstock", services.Addstock)
	//	r.POST("/login", services.Login)
	//r.Run()

}
