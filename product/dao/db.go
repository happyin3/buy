package dao

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Cateid     int64     `json:"cateid" gorm:"index"`
	Name       string    `json:"name"`
	Subtitle   string    `json:"subtitle"`
	Images     string    `json:"images"`
	Detail     string    `json:"detail"`
	Price      float64   `json:"price"`
	Stock      int64     `json:"stock"`
	Status     int64     `json:"status"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"upodate_time" gorm:"autoUpdateTime;index"`
}

func (Product) TableName() string {
	return "product"
}

var Database *gorm.DB

func ConnectDatabase() error {
	dsn := "root:MSqazwsx@tcp(127.0.0.1:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}
