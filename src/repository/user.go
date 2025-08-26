package repository

import (
	"database/sql"
	"fmt"
	"main/src/models"
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

func (repo UserRepository) GetUsers(slugOrName string) ([]models.User, error) {
	slugOrNameParseSqlLike := fmt.Sprintf("%%%s%%", slugOrName)

	rows, err := repo.db.Query("SELECT * FROM users WHERE name like ? OR slug like ?", slugOrNameParseSqlLike, slugOrNameParseSqlLike)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Slug,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}
