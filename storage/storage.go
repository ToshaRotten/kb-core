package storage

import (
	"database/sql"
	"fmt"
	"main/config"
)

type Storage struct {
	db *sql.DB
}

func New(config *config.Storage) (*Storage, error) {
	const op = "storage.New"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func GetLesson() {}

func AddLesson() {}

func DeleteLesson() {}

func UpdateLesson() {}

func GetGroup() {}

func AddGroup() {}

func DeleteGroup() {}

func UpdateGroup() {}

func GetProfessor() {}

func AddProfessor() {}

func DeleteProfessor() {}

func UpdateProfessor() {}
