package database

import (
	"fmt"
	"go_restfull_api/configs"
	"go_restfull_api/database/mysql"
	"go_restfull_api/models"
	"gorm.io/gorm"
)

/**
 * @author : Donald Trieu
 * @created : 9/1/21, Wednesday
**/

var MysqlDB *gorm.DB

func Init()  {
	c := configs.GetConfig()
	userName := c.GetString("db.mysql.username")
	pass := c.GetString("db.mysql.password")
	host := c.GetString("db.mysql.host")
	port := c.GetInt("db.mysql.port")
	dbName := c.GetString("db.mysql.dbname")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName, pass, host, port, dbName)
	MysqlDB = mysql.NewConnection(connStr)
}

func GetMysqlDB() *gorm.DB {
	return MysqlDB
}

func MigrateMysql() error {
	db := GetMysqlDB()
	err := db.AutoMigrate(&models.Camera{})
	return err
}