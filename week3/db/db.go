package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect
)

var (
	once        sync.Once
	mysqlClient *MySQLClient
)

const (
	username = "user"
	password = "password"
	dbname   = "feedback"
)

type MySQLClient struct {
	*gorm.DB
}

type PassengerFeedback struct {
	BookingCode string `gorm:"PRIMARY_KEY"`
	PassengerID int32  `gorm:"NOT NULL"`
	FeedBack    string `gorm:"NOT NULL"`
}

func init() {
	db := GetInstance()
	if !db.HasTable(&PassengerFeedback{}) {
		db.AutoMigrate(&PassengerFeedback{})
	}
}

// singleton once
func GetInstance() *MySQLClient {
	once.Do(func() {
		connStr := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", username, password, dbname)
		client, err := gorm.Open("mysql", connStr)
		log.Printf("connect to db: %v", connStr)
		if err != nil {
			log.Fatalf("cannot connect database: %v", err.Error())
		}
		client.LogMode(true)
		mysqlClient = &MySQLClient{client}
	})

	return mysqlClient
}
