package persistence

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetPool() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=username password=password dbname=dbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		var err error

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error getting sqlDB: %v", err)
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(time.Minute * 5)

		log.Println("Database connection pool initialized")
	})
	return db
}
