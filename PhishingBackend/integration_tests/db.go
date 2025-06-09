//go:build integration

package integration_tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log/slog"
	"os"
	"phishing_backend/internal/domain_model"
	"testing"
)

var gormDb *gorm.DB

func getUser(email string) *domain_model.User {
	user := &domain_model.User{}
	GetDb().Where("email = ?", email).First(user)
	return user
}

func createExam(t *testing.T, exam *domain_model.Exam) {
	result := GetDb().Create(exam)
	require.Nil(t, result.Error)
}

func GetDb() *gorm.DB {
	if gormDb == nil {
		gormDb = initGormAndDatabaseConnection()
	}
	return gormDb
}

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (d *dbConfig) getConnectionString() string {
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

func initGormAndDatabaseConnection() *gorm.DB {
	config := newDbConfig()
	connString := config.getConnectionString()
	slog.Info("Trying to connect to DB")
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Discard, // https://stackoverflow.com/a/55892341
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
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
	return db
}
