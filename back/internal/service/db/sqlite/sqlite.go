package sqlite

import (
	"database/sql"
	"log"

	"github.com/var10000/DnDHitPoints/back/internal/service/db"

	_ "github.com/mattn/go-sqlite3"
)

type userRepository struct {
	db *sql.DB
}

func CreateUserRepository() db.UserRepository {
	db, err := sql.Open("sqlite3", "dndbattles.db")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
		name TEXT NOT NULL UNIQUE
    )`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS characters (
      name TEXT NOT NULL,
      armor_type TEXT,
      initiative INTEGER,
      hits INTEGER,
      battle_id INTEGER,
      user_id INTEGER 
    )`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS rooms_users (
      room_id INTEGER, 
      user_id INTEGER
    )`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS battles_characters (
		battle_id INTEGER,
		character_id INTEGER,
		room_id INTEGER
    )`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	return &userRepository{db}
}

func (s *userRepository) AddUser(u db.UserDBModel) (db.UserDBModel, error) {
	stmt, err := s.db.Prepare(`INSERT into users (name) VALUES (?)`)
	if err != nil {
		return db.UserDBModel{}, err
	}
	_, err = stmt.Exec(u.Name)
	if err != nil {
		return db.UserDBModel{}, err
	}
	user, err := s.UserByName(u.Name)
	if err != nil {
		return db.UserDBModel{}, err
	}
	return user, nil
}

func (s *userRepository) DeleteUser(id int) error {
	stmt, err := s.db.Prepare(`DELETE from users where rowid = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *userRepository) UpdateUser(u db.UserDBModel) (db.UserDBModel, error) {
	stmt, err := s.db.Prepare(`UPDATE users set name = ? where rowid = ?`)
	if err != nil {
		return db.UserDBModel{}, err
	}
	_, err = stmt.Exec(u.Name, u.ID)
	if err != nil {
		return db.UserDBModel{}, err
	}
	return u, nil
}

func (s userRepository) User(id int) (db.UserDBModel, error) {
	stmt, err := s.db.Prepare(`select rowid, name from users where rowid = ?`)
	if err != nil {
		return db.UserDBModel{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return db.UserDBModel{}, err
	}
	var rowid int
	var name string
	row := stmt.QueryRow(id)
	err = row.Scan(&rowid, &name)
	if err != nil {
		return db.UserDBModel{}, err
	}
	u := db.UserDBModel{ID: rowid, Name: name}
	return u, nil
}

func (s userRepository) UserByName(userName string) (db.UserDBModel, error) {
	stmt, err := s.db.Prepare(`select rowid, name from users where name = ?`)
	if err != nil {
		return db.UserDBModel{}, err
	}
	_, err = stmt.Exec(userName)
	if err != nil {
		return db.UserDBModel{}, err
	}
	var rowid int
	var name string
	row := stmt.QueryRow(userName)
	err = row.Scan(&rowid, &name)
	if err != nil {
		return db.UserDBModel{}, err
	}
	u := db.UserDBModel{ID: rowid, Name: name}
	return u, nil
}
