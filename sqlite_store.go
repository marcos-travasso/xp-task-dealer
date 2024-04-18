package xp_task_dealer

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type SQLiteTasksStore struct {
	conn *gorm.DB
}

func InitDB() *SQLiteTasksStore {
	db, err := gorm.Open(sqlite.Open("./tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening sqlite file: %s", err)
	}

	db.AutoMigrate(&Task{})

	return &SQLiteTasksStore{conn: db}
}

func InitTestDB() *SQLiteTasksStore {
	os.Remove("./tasks_tests.db")
	db, err := gorm.Open(sqlite.Open("./tasks_tests.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening sqlite file: %s", err)
	}

	db.AutoMigrate(&Task{})

	return &SQLiteTasksStore{conn: db}
}

func (s *SQLiteTasksStore) SaveTask(task Task) error {
	result := s.conn.Save(&task)

	return result.Error
}

func (s *SQLiteTasksStore) GetTasks() ([]Task, error) {
	var tasks []Task

	result := s.conn.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (s *SQLiteTasksStore) GetTaskById(id string) (Task, error) {
	var task Task
	err := s.conn.First(&task, "id = ?", id).Error

	return task, err
}
