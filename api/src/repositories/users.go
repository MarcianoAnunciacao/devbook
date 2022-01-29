package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare(
		"insert into users (name, nick_name, email, password) values(?, ?, ?, ?)",
	)
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertedId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedId), nil
}
