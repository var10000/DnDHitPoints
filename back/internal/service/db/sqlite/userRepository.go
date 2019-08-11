package sqlite

import (
	"database/sql"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
)

type userRepository struct {
	db *sql.DB
}

const (
	CreateUserTableQuery = `CREATE TABLE IF NOT EXISTS users (user_name TEXT NOT NULL, id INTEGER PRIMARY KEY)`

	AddUserQuery     = `INSERT INTO users (user_name) VALUES (?)`
	DeleteUserQuery  = `DELETE FROM users WHERE id = ?`
	UpdateUserQuery  = `UPDATE users set name = ? where id = ?`
	GetByIDUserQuery = `SELECT user_name FROM users WHERE id = ?`
)

func (ur *userRepository) Add(u db.UserDBModel) (db.UserDBModel, error) {
	stmt, err := ur.db.Prepare(AddUserQuery)
	if err != nil {
		return db.UserDBModel{}, err
	}
	res, err := stmt.Exec(u.Name)
	if err != nil {
		return db.UserDBModel{}, err
	}
	u.ID, err = res.LastInsertId()
	return u, nil
}

func (ur *userRepository) Delete(id int64) error {
	stmt, err := ur.db.Prepare(DeleteUserQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Update(u db.UserDBModel) error {
	stmt, err := ur.db.Prepare(UpdateUserQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Name, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetByID(id int64) (db.UserDBModel, error) {
	stmt, err := ur.db.Prepare(GetByIDUserQuery)
	if err != nil {
		return db.UserDBModel{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return db.UserDBModel{}, err
	}
	var name string
	row := stmt.QueryRow(id)
	err = row.Scan(&name)
	if err != nil {
		return db.UserDBModel{}, err
	}
	u := db.UserDBModel{ID: id, Name: name}
	return u, nil
}
