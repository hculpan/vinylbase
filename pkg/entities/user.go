package entities

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/hculpan/vinylbase/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       int
	Username string
	Realname string
	password string
}

func NewUser(username, realname, password string) *User {
	result := &User{
		id:       -1,
		Username: username,
		Realname: realname,
	}
	result.SetPassword(password)
	return result
}

func CreateUserTable() error {
	_, err := db.Execute("CREATE TABLE users (username TEXT UNIQUE, realname TEXT, password TEXT)")
	if err != nil {
		log.Default().Printf("error creating users table: %s", err)
	} else {
		log.Default().Println("users table created")
	}
	return err
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
}

func FetchUser(username string) (*User, error) {
	var result *User = nil
	query := fmt.Sprintf("SELECT rowid, username, realname, password FROM users WHERE username='%s'", username)
	_, err := db.Query(query, db.QueryFunc(func(index int, rows *sql.Rows) error {
		user := User{}
		err := rows.Scan(&user.id, &user.Username, &user.Realname, &user.password)
		if err != nil {
			return err
		}

		result = &user

		return nil
	}))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *User) SaveUser() error {
	if u.password == "" {
		return errors.New("cannot save user with empty password")
	} else if u.Username == "" {
		return errors.New("cannot save user with empty username")
	}

	num, err := db.Execute(fmt.Sprintf("INSERT INTO users (username, realname, password) VALUES ('%s', '%s', '%s')", u.Username, u.Realname, u.password))
	if err != nil {
		return err
	} else if num != 1 {
		return fmt.Errorf("expected 1 row inserted, got %d", num)
	}

	return nil
}

func (u *User) DeleteUser() error {
	if u.id < 0 {
		return errors.New("cannot delete user with invalid rowid")
	}

	num, err := db.Execute(fmt.Sprintf("DELETE FROM users WHERE rowid=%d", u.id))
	if err != nil {
		return err
	} else if num != 1 {
		return fmt.Errorf("expected 1 row deleted, got %d", num)
	}

	return nil

}

func (u *User) Id() int {
	if u.id < 0 {
		user, err := FetchUser(u.Username)
		if err == nil && user != nil {
			u.id = user.id
		}
	}
	return u.id
}
