package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"select id, name, nick_name, email, created_at from users where name LIKE ? or nick_name LIKE ?",
		nameOrNick,
		nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchById(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick_name, email, created_at from users where id = ?",
		ID,
	)

	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick_name = ?, email = ?, where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.NickName, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository users) SearchByEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Follow(userID, followID uint64) error {
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

func (repository users) UnFollow(userID, followID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if _, err = statement.Exec(userID, followID); err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

func (repository users) SearchFollowersByUserID(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick_name, u.email, u.created_at
		from users u inner join followers f on u.id = f.follower_id
		where f.user_id = ?`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchUsersFollowedByAnUserID(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick_name, u.email, u.created_at
		from users u inner join followers f on u.id = f.user_id
		where f.user_id = ?`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchPasswordByUserId(userID uint64) (string, error) {
	line, err := repository.db.Query("select password from users where id = ?")
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository users) UpdatePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
