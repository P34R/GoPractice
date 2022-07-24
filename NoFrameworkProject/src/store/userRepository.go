package store

import (
	"GoTask1/src/model"
	"errors"
	"fmt"
)

type UserRepository struct {
	Store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if u.Password == "" {
		return errors.New("Incorrect password")
	}
	if err := u.EncryptPassword(); err != nil {
		return err
	}
	r.Store.db.QueryRow("INSERT INTO \"users\" (\"username\", \"password\",\"role\") VALUES ($1,$2,$3) ", u.Username, u.EncryptedPassword, u.Role)
	fmt.Println(u)
	return nil
}
func (r *UserRepository) Read(username string) (*model.User, error) {
	var u model.User
	err := r.Store.db.QueryRow("SELECT * FROM \"users\" WHERE \"username\"=$1", username).Scan(&u.Username, &u.EncryptedPassword, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *UserRepository) Delete(username string) error {
	_, err := r.Store.db.Exec("DELETE FROM \"users\" WHERE \"username\"=$1", username)
	if err != nil {
		return err
	}
	return nil
}
