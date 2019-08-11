package web

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type WebServer struct {
	UserResource   UserResource
	RoomResource   RoomResource
	BattleResource BattleResource
}

func (ws WebServer) SetRouters() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/rooms/{roomId}/characterList", ws.UserResource.getAllUsersCharacters)
		r.Post("/rooms/{roomId}/addUser/{userId}", ws.RoomResource.addRoomMembers)
		r.Delete("/rooms/{roomId}/removeUser/{userId}", ws.RoomResource.removeRoomMembers)
		r.Post("/battles/{battleId}/createBattle", ws.BattleResource.startFight)
		r.Get("/{roomId}/createBattle/{battleId}/characters", ws.BattleResource.getBattleCharacters)
	})
	return r
}

func (ws WebServer) StartAndServe(r chi.Router) {
	srv := &http.Server{Addr: ":10000", Handler: r}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Some errors have occurred, while running server: ", err)
	}
}
