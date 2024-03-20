package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Sqlite(url string) *Context {
	if err := os.MkdirAll(url, os.ModePerm); err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Silent),
		CreateBatchSize: 200,
	})
	if err != nil {
		log.Fatalln("error: database connection failed", err)
	}
	log.Println("database: sqlite: ", url)
	return &Context{DB: db}

}
