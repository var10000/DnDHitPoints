package sqlite

import (
	"database/sql"
	"log"

	"github.com/var10000/DnDHitPoints/back/internal/service/db"

	_ "github.com/mattn/go-sqlite3"
)

func initRepository(db *sql.DB, query string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func CreateRepositories() (db.UserRepository, db.CharacterRepository, db.RoomRepository, db.BattleRepository) {
	initQueries := []string{CreateBattleTableQuery, CreateRoomsTableQuery, CreateCharacterTableQuery, CreateUserTableQuery}

	db, err := sql.Open("sqlite3", "dndbattles.db")
	if err != nil {
		log.Fatal(err)
	}
	for _, query := range initQueries {
		err := initRepository(db, query)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &userRepository{db}, &characterRepository{db}, &roomRepository{db}, &battleRepository{db}
}




