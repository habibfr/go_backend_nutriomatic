package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang-template/models"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("No .env file found or failed to load")
	}

	// Load environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")

	var dbURI string
	if instanceConnectionName != "" {
		// Use Unix socket for Cloud SQL
		dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPassword, instanceConnectionName, dbName)
	} else {
		// Use TCP connection (default for local dev)
		if dbHost == "" {
			dbHost = "127.0.0.1"
		}
		if dbPort == "" {
			dbPort = "3306"
		}
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPassword, dbHost, dbPort, dbName)
	}

	// Custom logger
	newLogger := logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// Open DB connection
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Post-initialization action
	if err := UpdateProductIsShow(db); err != nil {
		log.Fatalf("Failed to update product_isShow: %v", err)
	}

	fmt.Println("Database connection successful")
	return db
}

func UpdateProductIsShow(db *gorm.DB) error {
	var products []models.Product
	currentDate := time.Now().Truncate(24 * time.Hour)

	err := db.Where("product_expShow = ?", currentDate).Find(&products).Error
	if err != nil {
		return err
	}

	for _, product := range products {
		product.PRODUCT_ISSHOW = 0
		if err := db.Save(&product).Error; err != nil {
			return err
		}
	}

	return nil
}
