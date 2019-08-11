package service

import (
	"github.com/var10000/DnDHitPoints/back/internal/service/db/sqlite"
)

// InitServices returns inited services
func InitServices() Application {
	ur, cr, rr, br := sqlite.CreateRepositories()
	return Application{
		ur: ur,
		cr: cr,
		rr: rr,
		br: br,
	}
}
