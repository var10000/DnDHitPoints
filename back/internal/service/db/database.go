package db

type UserDBModel struct {
	ID   int64
	Name string
}

type CharacterDBModel struct {
	ID         int64
	UserID     int64
	Name       string
	ArmorType  string
	Initiative int
	Hits       int
	BattleID   int64
}

type RoomDBModel struct {
	ID         int64
	UserID int64
}


type BattleDBModel struct {
	ID         int64
	CharacterID int64
	RoomID int64
}

type UserRepository interface {
	Add(u UserDBModel) (UserDBModel, error)
	Delete(id int64) error
	Update(u UserDBModel) error
	GetByID(id int64) (UserDBModel, error)
}

type CharacterRepository interface {
	Add(c CharacterDBModel) (CharacterDBModel, error)
	Delete(id int64) error
	Update(c CharacterDBModel) error
	GetByID(id int64) (CharacterDBModel, error)
	GetByUserID(userID int64) ([]CharacterDBModel, error)
}

type RoomRepository interface {
	Add(r RoomDBModel) (RoomDBModel, error)
	Delete(id int64) error
	Update(r RoomDBModel) error
	GetByID(id int64) (RoomDBModel, error)
}

type BattleRepository interface {
	Add(b BattleDBModel) (BattleDBModel, error)
	Delete(id int64) error
	Update(b BattleDBModel) error
	GetByID(id int64) (BattleDBModel, error)
}
