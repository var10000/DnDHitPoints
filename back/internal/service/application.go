package service

import (
	"fmt"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
	"log"
)

type Application struct {
	ur db.UserRepository
	br db.BattleRepository
	cr db.CharacterRepository
	rr db.RoomRepository
}

func (a *Application) Start() {
	a.testUser()
}

// ToRemove
func (a *Application) testUser() {
	u, err := a.ur.Add(db.UserDBModel{Name: "agafia2"})
	if err != nil {
		log.Fatal(err)
	}
	err = a.ur.Delete(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	user, err := a.ur.GetByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
