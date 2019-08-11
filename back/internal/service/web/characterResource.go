package web

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/var10000/DnDHitPoints/back/api"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
	"log"
	"net/http"
)

type CharacterResource struct {
	CharacterRepository db.CharacterRepository
}

func getMockedCharacters() api.CharacterListModel {
	return api.CharacterListModel{CharacterList: []api.CharacterModel{
		{
			Id:         1,
			Name:       "Good bitch",
			ArmorType:  "HEAVY",
			Hits:       20,
			Initiative: 20,
		},
		{
			Id:         2,
			Name:       "God bitch",
			ArmorType:  "LIGHT",
			Hits:       10,
			Initiative: 20,
		},
		{
			Id:         3,
			Name:       "Bad bitch",
			ArmorType:  "STEEL",
			Hits:       100,
			Initiative: 90,
		},
		{
			Id:         4,
			Name:       "Ugly bitch",
			ArmorType:  "WOODEN",
			Hits:       40,
			Initiative: -12,
		},
		{
			Id:         5,
			Name:       "Misha",
			ArmorType:  "DRAGON SKIN",
			Hits:       90,
			Initiative: -3,
		},
		{
			Id:         6,
			Name:       "Tvoya mama",
			ArmorType:  "LEATHER",
			Hits:       900,
			Initiative: 100,
		},
	},
	}
}

func (cr CharacterResource) getAllUsersCharacters(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r)
	roomId := chi.URLParam(r, "roomId")
	_ = roomId
	chars, err := cr.CharacterRepository.GetAllCharacters()
	marshalledUserList, _ := json.Marshal(chars)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marshalledUserList)
	if err != nil {
		log.Print("Wrong user model: ", err)
	}
}
