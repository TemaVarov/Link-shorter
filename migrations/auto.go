package main

import (
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/stat"
	"go/adv-demo/internal/user"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
		if err == nil {
			db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
			return
		}
		if i == 9 {
			panic(err)
		}
		log.Println(err.Error())
		time.Sleep(2 * time.Second)
	}
}
