package repository

import (
	"database/sql"
	"errors"
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
	defer rows.Close()

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
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}

func (repo UserRepository) GetUser(id uint64) (models.User, error) {

	rows, err := repo.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Slug,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

func (repo UserRepository) UpdateUser(userModel models.User) error {
	statement, err := repo.db.Prepare("UPDATE users SET name = ?, slug = ?, email = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userModel.Name, userModel.Slug, userModel.Email, userModel.UpdatedAt, userModel.ID)
	if err != nil {
		return err
	}

	return nil

}

func (repo UserRepository) DeleteUser(user_id uint64) error {
	statement, err := repo.db.Prepare("DELETE from users WHERE id = ?")
	if err != nil {
		return err
	}

	result, err := statement.Exec(user_id)
	if err != nil {
		return err
	}

	dataDeleted, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if dataDeleted == 0 {
		return errors.New("no user deleted")
	}

	return nil

}
