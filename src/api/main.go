package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	LoadEnv()

	var err error
	path := strings.Join([]string{os.Getenv("MYSQL_USER"), ":", os.Getenv("MYSQL_PASSWORD"), "@tcp(", os.Getenv("DB_HOST"), ":", os.Getenv("DB_PORT"), ")/", os.Getenv("MYSQL_DATABASE"), "?charset=utf8&parseTime=True&loc=Local"}, "")
	fmt.Println(path)

	db, err := gorm.Open("mysql", path)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Message{})
	// db.Create(&Message{Message: "Helloworld"})

	fmt.Println("Success connect gin_docker database")

	return db
}

// LoadEnv loading envfile
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}
