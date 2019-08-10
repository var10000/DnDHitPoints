package api

type CharacterModel struct {
	Img        string `json:"img"`
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	ArmorType  string `json:"armorType"`
	Hits       int32  `json:"hits"`
	Initiative int32  `json:"initiative"`
}

type CharacterListModel struct {
	UserList []CharacterModel `json:"userList"`
}
