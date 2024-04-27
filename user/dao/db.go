package dao

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string    `json:"username" gorm:"unique"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone" gorm:"unique"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime;index"`
}

func (User) TableName() string {
	return "user"
}

var Database *gorm.DB

func ConnectDatabase() error {
	dsn := "root:MSqazwsx@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}
