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

	database, err := sql.Open("sqlite3", "dndbattles.db")
	if err != nil {
		log.Fatal(err)
	}
	for _, query := range initQueries {
		err := initRepository(database, query)
		if err != nil {
			log.Fatal(err)
		}
	}
	characterRepository := characterRepository{database}
	characters := []db.CharacterDBModel{
		{ID: 1, UserID: 1, Name: "Naku", ArmorType: "LIGHT", Initiative: 6, Hits: 6, BattleID: 1},
		{ID: 2, UserID: 1, Name: "Lillian Songheart", ArmorType: "LIGHT", Initiative: 5, Hits: 8, BattleID: 1},
		{ID: 3, UserID: 2, Name: "Thorgrim", ArmorType: "HEAVY", Initiative: 8, Hits: 10, BattleID: 1},
		{ID: 4, UserID: 3, Name: "Gerald Immortal", ArmorType: "HEAVY", Initiative: 1, Hits: 16, BattleID: 1},
		{ID: 5, UserID: 4, Name: "Kojra", ArmorType: "HEAVY", Initiative: 6, Hits: 9, BattleID: 1},
		{ID: 6, UserID: 5, Name: "Orc The Banebreak Rider", ArmorType: "HEAVY", Initiative: 6, Hits: 16, BattleID: 1},
		{ID: 7, UserID: 6, Name: "Lidda Fireborn", ArmorType: "LIGHT", Initiative: 3, Hits: 7, BattleID: 1},
		{ID: 7, UserID: 6, Name: "TROLL KHAN", ArmorType: "HEAVY", Initiative: 3, Hits: 27, BattleID: 2},
	}
	for _, character := range characters {
		_, err = characterRepository.Add(character)
		if err != nil {
			log.Print("Imposible to add new character to character repository")
		}
	}
	return &userRepository{database}, &characterRepository, &roomRepository{database}, &battleRepository{database}
}
