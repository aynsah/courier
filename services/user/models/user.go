package models

import (
	"courier/pkg/token"
	"courier/services/user/database"
	"database/sql"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID     string `json:"uuid"`
	MSISDN   string `json:"msisdn"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credential struct {
	MSISDN   string `json:"msisdn"`
	Password string `json:"password"`
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}

	u.Password = string(bytes)

	return nil
}

func (u User) CheckPassword(pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))

	return err
}

func (u User) GenerateToken(JwtToken *token.JwtToken) error {
	var err error
	*JwtToken, err = token.Generate(u.UUID, u.MSISDN, u.Username)
	if err != nil {
		return err
	}

	return nil
}

func (u User) Save() (int, error) {
	if u.MSISDN == "" || u.Username == "" || u.Password == "" {
		return http.StatusBadRequest, errors.New("value cannot be nil")
	}

	if u.MSISDN[:2] != "62" {
		return http.StatusBadRequest, errors.New("invalid MSISDN")
	}

	db := database.DB

	err := u.HashPassword()
	if err != nil {
		return http.StatusInternalServerError, errors.New("hash password failed")
	}

	query := "INSERT INTO users (msisdn, username, password) VALUES (? , ?, ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.MSISDN, u.Username, u.Password)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func FindUserByMSISDN(MSISDN string, user *User) (int, error) {
	db := database.DB

	var err error

	query := `SELECT uuid, msisdn, username, password
			  FROM users
			  WHERE users.msisdn = ?`

	err = db.QueryRow(query, MSISDN).Scan(&user.UUID, &user.MSISDN, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		return http.StatusUnauthorized, errors.New("wrong username")
	} else if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func IsUserExists(MSISDN string, username string) (bool, error) {
	db := database.DB

	var err error
	var user User

	query := `SELECT uuid, msisdn, username, password
			  FROM users
			  WHERE users.msisdn = ?
			  OR users.username = ?`

	err = db.QueryRow(query, MSISDN, username).Scan(&user.UUID, &user.MSISDN, &user.Username, &user.Password)

	if err != sql.ErrNoRows {
		return true, err
	}

	return false, nil
}
