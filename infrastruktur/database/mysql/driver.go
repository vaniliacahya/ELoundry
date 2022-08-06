package mysql

import (
	"commerce-app/config"
	cartData "commerce-app/feature/cart/data"
	orderData "commerce-app/feature/order/data"
	productData "commerce-app/feature/product/data"
	userData "commerce-app/feature/user/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(userData.User{}, productData.Product{}, orderData.Order{}, cartData.Cart{})
}
