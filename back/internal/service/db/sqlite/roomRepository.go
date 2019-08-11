package sqlite

import (
	"database/sql"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
)

const (
	CreateRoomsTableQuery = `CREATE TABLE IF NOT EXISTS rooms(
		ID INTEGER,
		user_id INTEGER,
		PRIMARY KEY (ID, user_id)
	)`

	AddRoomQuery     = `INSERT INTO rooms (ID, user_id) VALUES (?, ?)`
	DeleteRoomQuery  = `DELETE FROM rooms WHERE ID = ?`
	UpdateRoomQuery  = `UPDATE rooms set user_id = ? WHERE ID = ?`
	GetByIDRoomQuery = `SELECT user_id FROM rooms WHERE ID = ?`
)

type roomRepository struct {
	db *sql.DB
}

func (rr *roomRepository) Add(r db.RoomDBModel) (db.RoomDBModel, error) {
	stmt, err := rr.db.Prepare(AddRoomQuery)
	if err != nil {
		return db.RoomDBModel{}, err
	}
	res, err := stmt.Exec(r.ID, r.UserID)
	if err != nil {
		return db.RoomDBModel{}, err
	}
	r.ID, err = res.LastInsertId()
	return r, nil
}

func (rr *roomRepository) Delete(id int64) error {
	stmt, err := rr.db.Prepare(DeleteRoomQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (rr *roomRepository) Update(r db.RoomDBModel) error {
	stmt, err := rr.db.Prepare(UpdateRoomQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(r.UserID, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func (rr roomRepository) GetByID(id int64) (db.RoomDBModel, error) {
	stmt, err := rr.db.Prepare(GetByIDRoomQuery)
	if err != nil {
		return db.RoomDBModel{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return db.RoomDBModel{}, err
	}
	var userID int64
	row := stmt.QueryRow(id)
	err = row.Scan(&userID)
	if err != nil {
		return db.RoomDBModel{}, err
	}
	u := db.RoomDBModel{ID: id, UserID: userID}
	return u, nil
}
