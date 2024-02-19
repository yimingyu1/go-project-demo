package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)

var DB *gorm.DB

func Init() *gorm.DB {
	return connectDB(
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.name"),
	)
}

func connectDB(user, password, ip, dbName string) *gorm.DB {
	config := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		password,
		ip,
		dbName,
		true,
		//"Asia/Shanghai"),
		"Local",
	)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		println("fuck")
		//log.Errorf("Database connection failed. Database name: %s, err: %+v", name, err)
		panic(err)
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	// set for db connection
	db.LogMode(viper.GetBool("mysql.show_log"))
	// To set the maximum number of open connections, replace with 0 to indicate unlimited.
	// Setting the maximum number of connections can avoid too high concurrency
	// leading to too many connection errors when connecting to mysql.
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open_conn"))
	// Used to set the number of idle connections. When the number of idle connections
	// is set, the opened connection can be placed in the pool for the next use.
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conn"))
	db.DB().SetConnMaxLifetime(time.Minute * viper.GetDuration("mysql.conn_max_life_time"))

	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}

func MigrateDB() {
	DB.AutoMigrate(&ListModel{})
}
