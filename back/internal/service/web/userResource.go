package web

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/var10000/DnDHitPoints/back/api"
	"log"
	"net/http"
)

type UserResource struct {
}

func (ur UserResource) getAllUsersCharacters(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	_ = roomId
	w.WriteHeader(http.StatusOK)
	var userList = api.CharacterListModel{UserList: []api.CharacterModel{
		{
			Img:        "https://en.wikipedia.org/wiki/Orc#/media/File:Orc_mask_by_GrimZombie.jpg",
			Id:         1,
			Name:       "Good bitch",
			ArmorType:  "HEAVY",
			Hits:       20,
			Initiative: 20,
		},
		{
			Img:        "https://en.wikipedia.org/wiki/Orc#/media/File:Orc_mask_by_GrimZombie.jpg",
			Id:         2,
			Name:       "God bitch",
			ArmorType:  "LIGHT",
			Hits:       10,
			Initiative: 20,
		},
		{
			Img:        "https://en.wikipedia.org/wiki/Orc#/media/File:Orc_mask_by_GrimZombie.jpg",
			Id:         3,
			Name:       "Bad bitch",
			ArmorType:  "STEEL",
			Hits:       100,
			Initiative: 90,
		},
		{
			Img:        "https://en.wikipedia.org/wiki/Orc#/media/File:Orc_mask_by_GrimZombie.jpg",
			Id:         4,
			Name:       "Ugly bitch",
			ArmorType:  "WOODEN",
			Hits:       40,
			Initiative: -12,
		},
		{
			Img:        "https://en.wikipedia.org/wiki/Orc#/media/File:Orc_mask_by_GrimZombie.jpg",
			Id:         5,
			Name:       "Misha",
			ArmorType:  "DRAGON SKIN",
			Hits:       90,
			Initiative: -3,
		},
		{
			Img:        "https://en.wikipedia.org/wiki/Lizzo#/media/File:Lizzo_2019_MTV.png",
			Id:         6,
			Name:       "Tvoya mama",
			ArmorType:  "LEATHER",
			Hits:       900,
			Initiative: 100,
		},
	},
	}
	marshalledUserList, _ := json.Marshal(userList)
	_, err := w.Write(marshalledUserList)
	if err != nil {
		log.Print("Wrong user model: ", err)
	}
}
