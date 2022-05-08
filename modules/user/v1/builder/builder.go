package builder

import (
	userRepo "gin-starter/modules/user/v1/repository"
	"gin-starter/modules/user/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
)

// BuildUserHandler build user handlers
// starting from handler down to repository or tool.
func BuildUserHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool) {
	// Repository
	ur := userRepo.NewUserRepository(db)

	// Service
	uf := service.NewUserFinder(cfg, ur)
	uu := service.NewUserUpdater(cfg, ur)

	app.UserFinderHTTPHandler(cfg, router, uf)
	app.UserUpdaterHTTPHandler(cfg, router, uu, uf)
}
