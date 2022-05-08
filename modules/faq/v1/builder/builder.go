package builder

import (
	"gin-starter/app"
	"gin-starter/config"
	faqRepo "gin-starter/modules/faq/v1/repository"
	faq "gin-starter/modules/faq/v1/service"
	"github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BuildFaqHandler build user handlers
// starting from handler down to repository or tool.
func BuildFaqHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool) {
	// Repository
	faqRp := faqRepo.NewFaqRepository(db)

	// Service
	nf := faq.NewFaqFinder(
		cfg,
		faqRp,
	)

	app.FaqFinderHTTPHandler(cfg, router, nf)
}
