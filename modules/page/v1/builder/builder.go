package builder

import (
	"gin-starter/app"
	"gin-starter/config"
	pageRepo "gin-starter/modules/page/v1/repository"
	page "gin-starter/modules/page/v1/service"
	"github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BuildPageHandler build user handlers
// starting from handler down to repository or tool.
func BuildPageHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool) {
	// Repository
	pageRp := pageRepo.NewPageRepository(db)

	// Service
	nf := page.NewPageFinder(
		cfg,
		pageRp,
	)

	app.PageFinderHTTPHandler(cfg, router, nf)
}
