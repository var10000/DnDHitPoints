package service

import (
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
	"github.com/var10000/DnDHitPoints/back/internal/service/db/sqlite"
)

// InitServices returns inited services
func InitServices() (db.UserRepository, db.CharacterRepository, db.RoomRepository, db.BattleRepository){
	userRepo, charRepo, roomRepo, battleRepo := sqlite.CreateRepositories()
	return userRepo, charRepo, roomRepo, battleRepo
}
