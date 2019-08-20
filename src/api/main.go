package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Message struct {
	Message string
}

func main() {
	db := dbConnect()
	defer db.Close()
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":3001")
}

func dbConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(godockerDB)/godocker")

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Message{})
	db.Create(&Message{Message: "Helloworld"})

	return db
}
