package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbConnect *sql.DB

func main() {
	DbConnect, err := sql.Open("mysql", "root:@tcp(db:3306)/godocker")
	if err != nil {
		panic(err.Error())
	}
	DbConnect.Close()
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":3001")
}
