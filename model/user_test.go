package model

import (
	"fmt"
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mockDB *gorm.DB

func TestMain(m *testing.M) {
	dsn := "root:qq3200334@tcp(127.0.0.1:3306)/moft?charset=utf8mb4&parseTime=True"
	var err error
	mockDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("connect database failed: %v", err)
		return
	}
	os.Exit(m.Run())
}

func TestListUser(t *testing.T) {
	users, err := ListUser(mockDB, map[string]interface{}{
		"user_name": "liuchengshun",
		"password":  "liuchengshun",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("users", users)
}
