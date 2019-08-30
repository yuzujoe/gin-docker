package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
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
	Load_Env()
	// db, err := gorm.Open("mysql", "root:root@tcp(godockerDB)/godocker")

	db, err := gorm.Open("mysql", os.Getenv("MYSQL_ROOT_USER")+":"+os.Getenv("MYSQL_ROOT_PASSWORD")+"@tcp(godockerDB)/"+os.Getenv("MYSQL_DATABASE"))

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Message{})
	db.Create(&Message{Message: "Helloworld"})

	return db
}

func Load_Env() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}
