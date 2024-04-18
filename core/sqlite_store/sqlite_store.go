package sqlite_store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"xp-task-dealer/core/models"
)

type SQLiteStore struct {
	conn *gorm.DB
}

func Init(dbDir string) *SQLiteStore {
	db, err := gorm.Open(sqlite.Open(dbDir), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening sqlite file: %s", err)
	}

	db.AutoMigrate(&models.Task{}, &models.Developer{})

	return &SQLiteStore{conn: db}
}

func InitTest() *SQLiteStore {
	os.Remove("./tasks_tests.db")
	return Init("./tasks_tests.db")
}
