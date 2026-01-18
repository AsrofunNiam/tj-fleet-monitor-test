package app

import (
	"fmt" // Tambahkan fmt untuk merapikan DSN
	"log"
	"os"
	"time"

	"github.com/AsrofunNiam/tj-fleet-monitor-test/model/domain"
	"gorm.io/driver/postgres" // Ganti ini
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// Format DSN untuk PostgreSQL berbeda dengan MySQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, db, port)

	// Gunakan postgres.Open bukan mysql.Open
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto Migrate tetap sama
	err = database.AutoMigrate(&domain.VehicleLocation{})
	if err != nil {
		panic("failed to auto migrate schema")
	}

	return database
}
