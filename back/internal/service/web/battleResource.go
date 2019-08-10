package web

import (
	"github.com/go-chi/chi"
	"net/http"
)

type BattleResource struct {
}

func (br BattleResource) addCharacter(w http.ResponseWriter, r *http.Request) {
	battleId := chi.URLParam(r, "battleId")
	characterId := chi.URLParam(r, "characterId")
	_ = battleId
	_ = characterId
}

func (br BattleResource) removeCharacter(w http.ResponseWriter, r *http.Request) {
	battleId := chi.URLParam(r, "battleId")
	characterId := chi.URLParam(r, "characterId")
	_ = battleId
	_ = characterId
}

func (br BattleResource) getBattleCharacters(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	battleId := chi.URLParam(r, "battleId")
	characterId := chi.URLParam(r, "characterId")
	_ = roomId
	_ = battleId
	_ = characterId
}
