package persistance

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func (d *dbConfig) getPostgresConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", d.user, d.password, d.host, d.port, d.dbname)
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
	err := migrateDatabaseSchema()
	if err != nil {
		slog.Error("Database schema could not be migrated", "err", err)
		panic(err)
	}
	//createTables()
}

func initGormAndDatabaseConnection() {
	config := newDbConfig()
	connString := config.getPostgresConnString()
	slog.Info("Trying to connect to DB")

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

//func createTables() {
//	if err := db.AutoMigrate(&domain.User{}, &domain.LessonCompletion{}); err != nil {
//		slog.Error("Could not create table", "error", err)
//	}
//}
