package repository

import (
	"database/sql"
	"main/src/router/models"
)

type UserRepository struct {
	db *sql.DB
}

// constructor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo UserRepository) Create(userModel models.User) (int, error) {
	statment, err := repo.db.Prepare("INSERT INTO users (name, slug, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer statment.Close()

	result, err := statment.Exec(userModel.Name, userModel.Slug, userModel.Email, userModel.Password)
	if err != nil {
		return 0, err
	}

	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastIdInserted), nil
}
