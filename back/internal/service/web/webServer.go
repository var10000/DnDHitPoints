package web

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"math/rand"
)

type WebServer struct {
	CharacterResource CharacterResource
	RoomResource      RoomResource
	BattleResource    BattleResource
}

func getColor(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r)
	rand.Seed(42)
	answers := []string{
		"blue",
		"green",
		"red",
	}
	col := answers[rand.Intn(len(answers))]
	marshalledUserList, _ := json.Marshal(col)
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

func (ws WebServer) SetRouters() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/color", )
		r.Get("/rooms/{roomId}/characterList", ws.CharacterResource.getAllUsersCharacters)
		r.Post("/rooms/{roomId}/getOrder", ws.CharacterResource.getSortedForRoom)
		r.Options("/rooms/{roomId}/characterList", ws.CharacterResource.getAllUsersCharacters)
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
