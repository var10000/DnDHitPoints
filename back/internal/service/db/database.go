package db

type UserDBModel struct {
	ID   int
	Name string
}

type CharacterDBModel struct {
	ID         int
	Name       string
	ArmorType  string
	Initiative int
	BattleID   int
}

type UserRepository interface {
	AddUser(u UserDBModel) (UserDBModel, error)
	DeleteUser(id int) error
	UpdateUser(u UserDBModel) (UserDBModel, error)
	User(id int) (UserDBModel, error)
}
