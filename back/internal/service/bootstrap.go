package service

import (
	"github.com/var10000/DnDHitPoints/back/internal/service/db"
	"github.com/var10000/DnDHitPoints/back/internal/service/db/sqlite"
)

// InitServices returns inited services
func InitServices() db.UserRepository {
	ur := sqlite.CreateUserRepository()
	return ur
}
