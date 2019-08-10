package sqlite

import (
	"database/sql"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
)

const (
	CreateBattleTableQuery = `CREATE TABLE IF NOT EXISTS battles (
		character_id INTEGER,
		room_id INTEGER
    )`

	AddBattleQuery = `INSERT INTO battles (name) VALUES (?)`
	DeleteBattleQuery = `DELETE FROM battles WHERE rowid = ?`
	UpdateBattleQuery = `UPDATE battles set character_id = ?, room_id = ? where rowid = ?`
	GetByIDBattleQuery = `SELECT rowid, name FROM users WHERE rowid = ?`
)


type battleRepository struct {
	db *sql.DB
}
func (br *battleRepository) Add(b db.BattleDBModel) (db.BattleDBModel, error) {
	stmt, err := br.db.Prepare(AddBattleQuery)
	if err != nil {
		return db.BattleDBModel{}, err
	}
	res, err := stmt.Exec(b.CharacterID, b.RoomID)
	if err != nil {
		return db.BattleDBModel{}, err
	}
	b.ID, err = res.LastInsertId()
	return b, nil
}

func (br *battleRepository) Delete(id int64) error {
	stmt, err := br.db.Prepare(DeleteBattleQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (br *battleRepository) Update(b db.BattleDBModel) error {
	stmt, err := br.db.Prepare(UpdateBattleQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(b.CharacterID, b.RoomID)
	if err != nil {
		return err
	}
	return nil
}

func (br *battleRepository) GetByID(id int64) (db.BattleDBModel, error) {
	stmt, err := br.db.Prepare(GetByIDBattleQuery)
	if err != nil {
		return db.BattleDBModel{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return db.BattleDBModel{}, err
	}
	var roomID int64
	var CharacterID int64
	row := stmt.QueryRow(id)
	err = row.Scan(&CharacterID, &roomID)
	if err != nil {
		return db.BattleDBModel{}, err
	}
	b := db.BattleDBModel{
		ID: id,
		RoomID: roomID,
		CharacterID:CharacterID,
	}
	return b, nil
}

