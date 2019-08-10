package service

import (
	"github.com/var10000/DnDHitPoints/back/internal/service/db/sqlite"
)

// InitServices returns inited services
func InitServices() {
	sqlite.CreateRepositories()
}
