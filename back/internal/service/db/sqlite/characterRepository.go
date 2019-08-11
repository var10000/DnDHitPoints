package sqlite

import (
	"database/sql"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
)

const (
	CreateCharacterTableQuery = `CREATE TABLE IF NOT EXISTS characters (
      character_id INTEGER PRIMARY KEY,
      character_name TEXT NOT NULL,
      armor_type TEXT,
      initiative INTEGER,
      hits INTEGER,
      battle_id INTEGER,
      user_id INTEGER 
    )`

	AddCharacterQuery = `INSERT into characters (character_name, armor_type, initiative, hits, battle_id, user_id) VALUES (?, ?, ?, ?, ?, ?)`
	DeleteCharacterQuery = `DELETE FROM characters WHERE character_id = ?`
	UpdateCharacterQuery = `UPDATE rooms set name = ?, armor_type = ?, initiative = ?, hits = ?, battle_id = ?, user_id = ? WHERE rowid = ?`
	GetByIDCharacterQuery = `SELECT name, armor_type, initiative, hits, battle_id, user_id FROM characters WHERE rowid = ?`
	GetByUserIDCharacterQuery = `SELECT character_id, character_name, armor_type, initiative, hits, battle_id FROM characters WHERE user_id = ?`
)

type characterRepository struct {
	db *sql.DB
}


func (cr *characterRepository) Add(c db.CharacterDBModel) (db.CharacterDBModel, error) {
	stmt, err := cr.db.Prepare(AddCharacterQuery)
	if err != nil {
		return db.CharacterDBModel{}, err
	}
	res, err := stmt.Exec(c.Name, c.ArmorType, c.Initiative, c.Hits, c.BattleID, c.UserID)
	if err != nil {
		return db.CharacterDBModel{}, err
	}
	c.ID, err = res.LastInsertId()
	if err != nil {
		return db.CharacterDBModel{}, err
	}
	return c, nil
}

func (cr *characterRepository) Delete(id int64) error {
	stmt, err := cr.db.Prepare(DeleteCharacterQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (cr *characterRepository) Update(c db.CharacterDBModel) error {
	stmt, err := cr.db.Prepare(UpdateCharacterQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.Name, c.ArmorType, c.Initiative, c.Hits, c.BattleID, c.UserID, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cr *characterRepository) GetByID(id int64) (db.CharacterDBModel, error) {
	stmt, err := cr.db.Prepare(GetByIDCharacterQuery)
	if err != nil {
		return db.CharacterDBModel{}, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return db.CharacterDBModel{}, err
	}
	var name, armorType  string
	var initiative, hits int
	var battleID, userID int64
	row := stmt.QueryRow(id)
	err = row.Scan(&name, &armorType, &initiative, hits, battleID, userID)
	if err != nil {
		return db.CharacterDBModel{}, err
	}
	c := db.CharacterDBModel{
		ID: id,
		Name: name,
		Initiative: int(initiative),
		Hits: int(hits),
		BattleID: battleID,
		UserID: userID,
	}
	return c, nil
}


func (cr *characterRepository) GetByUserID(id int64) ([]db.CharacterDBModel, error) {
	// TODO returns one, make return all
	stmt, err := cr.db.Prepare(GetByUserIDCharacterQuery)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return nil, err
	}
	var name, armorType string
	var initiative, hits int
	var characterID, battleID int64
	row := stmt.QueryRow(id)
	err = row.Scan(&characterID, &name, &armorType, &initiative, &hits, &battleID)
	if err != nil {
		return nil, err
	}
	c := db.CharacterDBModel{
		ID: characterID,
		Name: name,
		Initiative: int(initiative),
		Hits: int(hits),
		BattleID: battleID,
		UserID: id,
	}
	return []db.CharacterDBModel{c}, nil
}
