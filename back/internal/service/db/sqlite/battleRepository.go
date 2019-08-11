package sqlite

import (
	"database/sql"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
)

const (
	CreateBattleTableQuery = `CREATE TABLE IF NOT EXISTS battles (
		battle_id INTEGER,
		character_id INTEGER,
		PRIMARY KEY (battle_id, character_id)
    )`

	AddBattleQuery     = `INSERT INTO battles (battle_id, character_id) VALUES (?, ?)`
	DeleteBattleQuery  = `DELETE FROM battles WHERE character_id = ?`
	UpdateBattleQuery  = `UPDATE battles SET character_id = ? where battle_id = ?`
	GetByIDBattleQuery = `SELECT (battle_id, character_id) name FROM users WHERE battle_id = ?`
)

type battleRepository struct {
	db *sql.DB
}

func (br *battleRepository) Add(b db.BattleDBModel) (db.BattleDBModel, error) {
	stmt, err := br.db.Prepare(AddBattleQuery)
	if err != nil {
		return db.BattleDBModel{}, err
	}
	res, err := stmt.Exec(b.ID, b.CharacterID)
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
	_, err = stmt.Exec(b.CharacterID, b.ID)
	if err != nil {
		return err
	}
	return nil
}

func (br *battleRepository) GetByID(id int64) ([]db.BattleDBModel, error) {
	stmt, err := br.db.Prepare(GetByIDBattleQuery)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return nil, err
	}
	var roomID int64
	var CharacterID int64
	row := stmt.QueryRow(id)
	err = row.Scan(&CharacterID, &roomID)
	if err != nil {
		return nil, err
	}
	//TODO: возвращаем одну, потом уже сделаем несколько
	b := []db.BattleDBModel{
		{
			ID:          id,
			CharacterID: CharacterID,
		},
	}
	return b, nil
}
