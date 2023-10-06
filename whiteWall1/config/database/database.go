package database

import (
	"fmt"
	"log"
	"whiteWall/config/config"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	user := config.Config.GetString("database.user")
	pass := config.Config.GetString("database.pass")
	port := config.Config.GetString("database.DBport")
	host := config.Config.GetString("database.host")
	name := config.Config.GetString("database.DBname")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Database connect failed: ", err)
	}

	// 自动建表（已有不会覆盖）
	err = autoMigrate(db)
	if err != nil {
		log.Fatal("Database migrate failed: ", err)
	}

	DB = db
}
