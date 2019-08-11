package web

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/var10000/DnDHitPoints/back/api"
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
	"log"
	"net/http"
	"strconv"
)

type BattleResource struct {
	battleRepository db.BattleRepository
}

func (br BattleResource) startFight(w http.ResponseWriter, r *http.Request) {
	battleId, _ := strconv.Atoi(chi.URLParam(r, "battleId"))
	decoder := json.NewDecoder(r.Body)
	var characterList api.CharacterListModel
	err := decoder.Decode(&characterList)
	if err != nil {
		panic(err)
	}
	for _, character := range characterList.CharacterList {
		b := db.BattleDBModel{
			ID:          int64(battleId),
			CharacterID: int64(character.Id),
		}
		_, err := db.BattleRepository.Add(battleId, b)
		if err != nil {
			log.Fatal("Request for start a fight was failed with error: ", err)
		}
	}
}

func (br BattleResource) getBattleCharacters(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	battleId := chi.URLParam(r, "battleId")
	_ = roomId
	_ = battleId
}
