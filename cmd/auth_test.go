package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "ggooo.povarov@gmail.com",
		Password: "$2a$10$nfgTxI6T19eR0h0BRkWjo.s3zZmmOec8teHQpXgVvnUKsxUlZJwRa",
		NickName: "Вася",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "ggooo.povarov@gmail.com").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	// Prepare
	db := initDB()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "ggooo.povarov@gmail.com",
		Password: "111112222",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var tokenJson auth.LoginResponce
	err = json.Unmarshal(body, &tokenJson)
	if err != nil {
		t.Fatal(err)
	}
	if tokenJson.Token == "" {
		t.Fatal("Token is empty")
	}
	removeData(db)
}

func TestLoginFail(t *testing.T) {
	// Prepare
	db := initDB()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "ggooo.povarov@gmail.com",
		Password: "11111222",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}
	removeData(db)
}
