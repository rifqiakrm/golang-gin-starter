package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
	pageRepo "gin-starter/modules/page/v1/repository"
	page "gin-starter/modules/page/v1/service"
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
