package web

import (
	"github.com/go-chi/chi"
	"net/http"
)

type RoomResource struct {
}

func (rr RoomResource) addRoomMembers(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	userId := chi.URLParam(r, "userId")
	_ = roomId
	_ = userId
}

func (rr RoomResource) removeRoomMembers(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	userId := chi.URLParam(r, "userId")
	_ = roomId
	_ = userId
}
