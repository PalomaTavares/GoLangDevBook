package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// repositoru of users
type Users struct {
	db *sql.DB
}

// creates user repository
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, senha) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil
}

func (repository Users) Get(nameORNick string) ([]models.User, error) {
	nameORNick = fmt.Sprintf("%%%s%%", nameORNick) //%

	lines, error := repository.db.Query(
		"select id, name, nick, email, createdIn from users where name LIKE ? or nick LIKE ?",
		nameORNick, nameORNick,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository Users) GetByID(ID uint64) (models.User, error) {
	lines, error := repository.db.Query(
		"select id, name, nick, email, createdIn from users where id = ?",
		ID,
	)
	if error != nil {
		return models.User{}, error
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, error := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error
	}
	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, error := repository.db.Prepare("delete from users where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(ID); error != nil {
		return error
	}
	return nil
}

func (repository Users) GetByEmail(email string) (models.User, error) {
	line, error := repository.db.Query("select id, senha from users where email = ?", email)
	defer line.Close()

	var user models.User

	if line.Next() {
		if error = line.Scan(
			&user.ID,
			&user.Password,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}
