package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "db"
	dbName := "employee_db"
	dbUser := "admin"
	dbPass := "admin123"
	dbPort := "5432"
	sslmode := "disable"
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslmode,
	)

	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("failed to open sql connection:", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("failed to create migrate driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://employee-database/db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("failed to create migrate instance:", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("failed to apply migrations:", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect with GORM:", err)
	}

	DB = gormDB
	log.Println("✅ Database initialized and migrated successfully")
}
