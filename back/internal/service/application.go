package service

import "github.com/var10000/DnDHitPoints/back/internal/service/db"

type Application struct {
	ur db.UserRepository
	br db.BattleRepository
	cr db.CharacterRepository
	rr db.RoomRepository
}

func (a *Application) Start() {

}
