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
	statement, err := repo.db.Prepare("INSERT INTO users (name, slug, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(userModel.Name, userModel.Slug, userModel.Email, userModel.Password)
	if err != nil {
		return 0, err
	}

	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastIdInserted), nil
}
