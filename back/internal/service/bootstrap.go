package service

import (
	"github.com/var10000/DnDHitPoints/back/internal/service/db/sqlite"
	"github.com/var10000/DnDHitPoints/back/internal/service/web"
)

// InitServices returns inited services
func InitServices() Application {
	ur, cr, rr, br := sqlite.CreateRepositories()
	ws := web.WebServer{CharacterResource: web.CharacterResource{CharacterRepository: cr}}
	return Application{
		ur:   ur,
		cr:   cr,
		rr:   rr,
		br:   br,
		serv: ws,
	}
}
