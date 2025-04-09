package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"phishing_backend/internal/domain/model"

	"log/slog"
	"os"
)

var db *gorm.DB

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (d *dbConfig) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.dbname)
}

func newDbConfig() *dbConfig {
	return &dbConfig{
		host:     os.Getenv("PHBA_DB_HOST"),
		port:     os.Getenv("PHBA_DB_PORT"),
		user:     os.Getenv("PHBA_DB_USER"),
		password: os.Getenv("PHBA_DB_PASSWORD"),
		dbname:   os.Getenv("PHBA_DB_NAME"),
	}
}

func init() {
	initGormAndDatabaseConnection()
	createTables()
}

func initGormAndDatabaseConnection() {
	config := newDbConfig()
	connString := config.getConnectionString()
	slog.Info("Trying to connect to DB", "connectionString", connString)

	var err error
	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Discard, // https://stackoverflow.com/a/55892341
	})
	if err != nil {
		slog.Error("Could not connect to db", "error", err)
		os.Exit(1)
	}
	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Could not sqlDB of gorm", "error", err)
		os.Exit(1)
	}
	err = sqlDB.Ping()
	if err != nil {
		slog.Error("Could not ping db", "error", err)
		os.Exit(1)
	}
	slog.Info("Connection to db successful")
}

func createTables() {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		slog.Error("Could not create table", "error", err)
	}
}
