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
	a.testCharacter()
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



func (a *Application) testCharacter() {
	c, err := a.cr.Add(db.CharacterDBModel{Name: "aa", UserID: 12, ArmorType: "yuiooiuy", Initiative: 4, Hits: 5, BattleID: 45})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	char, err := a.cr.GetByUserID(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(char)
	err = a.cr.Delete(1)
	if err != nil {
		log.Fatal(err)
	}
}
