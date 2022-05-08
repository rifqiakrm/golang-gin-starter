package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
	faqRepo "gin-starter/modules/faq/v1/repository"
	faq "gin-starter/modules/faq/v1/service"
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
