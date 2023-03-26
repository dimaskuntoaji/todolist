package db

import (
	"fmt"

	"todolist/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("MYSQL_USER", "root"),
		config.GetEnv("MYSQL_PASSWORD", ""),
		config.GetEnv("MYSQL_HOST", "localhost"),
		config.GetEnv("MYSQL_DBNAME", "todo"),
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
