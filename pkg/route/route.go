package route

import (
	"github.com/gin-contrib/pprof"
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/service"
	"gitlab.com/merakilab9/meracrawler/kayle/conf"
	"gitlab.com/merakilab9/meracrawler/kayle/pkg/repo/pg"

	handlerTransform "gitlab.com/merakilab9/meracrawler/kayle/pkg/handler"
	serviceTransform "gitlab.com/merakilab9/meracrawler/kayle/pkg/service"
)

type Service struct {
	*service.BaseApp
}

func NewService() *Service {

	s := &Service{
		service.NewApp("kayle Service", "v1.0"),
	}
	db := s.GetDB()

	if !conf.LoadEnv().DbDebugEnable {
		db = db.Debug()
	}

	repoPG := pg.NewPGRepo(db)
	var transformService = serviceTransform.NewTransformService(repoPG)
	transformHandle := handlerTransform.NewTransformHandlers(transformService)
	migrateHandle := handlerTransform.NewMigrationHandler(db)

	pprof.Register(s.Router)

	v1Api := s.Router.Group("/api/v1")
	v1Api.POST("/media/pre-upload", ginext.WrapHandler(transformHandle.Transform))

	v1Api.POST("/internal/migrate", migrateHandle.Migrate)

	return s
}
