# gin-docker

## あとで使用
db.AutoMigrate(&Message{})
	db.Create(&Message{Message: "Helloworld"})
type Message struct {
	gorm.Model
	Message string
}
