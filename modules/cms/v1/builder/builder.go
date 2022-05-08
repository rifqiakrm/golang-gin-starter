package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
	"gin-starter/modules/cms/v1/repository"
	"gin-starter/modules/cms/v1/service"
	notificationRepo "gin-starter/modules/notification/v1/repository"
	notification "gin-starter/modules/notification/v1/service"
)

// BuildCMSHandler build user handlers
// starting from handler down to repository or tool.
func BuildCMSHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool) {
	// Repository
	userRepo := repository.NewUserRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	notifRepo := notificationRepo.NewNotificationRepository(db)
	pageRepo := repository.NewPageRepository(db)
	faqRepo := repository.NewFaqRepository(db)

	// Service
	nc := notification.NewNotificationCreator(cfg, notifRepo)

	cf := service.NewCMSFinder(
		cfg,
		userRepo,
		roleRepo,
		faqRepo,
		pageRepo,
	)

	cc := service.NewCMSCreator(
		cfg,
		userRepo,
		userRoleRepo,
		nc,
		faqRepo,
		pageRepo,
	)
	cu := service.NewCMSUpdater(
		cfg,
		userRepo,
		userRoleRepo,
		nc,
		faqRepo,
		pageRepo,
	)
	cd := service.NewCMSDeleter(
		cfg,
		userRepo,
		faqRepo,
		pageRepo,
	)

	app.CMSFinderHTTPHandler(cfg, router, cf)
	app.CMSCreatorHTTPHandler(cfg, router, cc, cf)
	app.CMSUpdaterHTTPHandler(cfg, router, cu, cf)
	app.CMSDeleterHTTPHandler(cfg, router, cd)
}
