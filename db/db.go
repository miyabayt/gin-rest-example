package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/miyabayt/gin-rest/config"
	"github.com/miyabayt/gin-rest/logger"
)

var (
	DB *gorm.DB
)

func init() {
	dialectName, _ := config.String("db.dialect")
	connectionString, _ := config.String("db.connection_string")

	db, err := gorm.Open(dialectName, connectionString)
	if err != nil {
		logger.Errorf("failed to open connection. dialect=%s, connectionString=%s", dialectName, connectionString)
		panic(err)
	}

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(2)  // TODO
	db.DB().SetMaxOpenConns(10) // TODO

	db.LogMode(true) // TODO
	// db.SetLogger(Logger) // TODO format, implement gorm.Logger

	DB = &db
}
