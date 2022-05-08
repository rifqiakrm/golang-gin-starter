package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
	"gin-starter/modules/notification/v1/pubsub/handler"
	"gin-starter/modules/notification/v1/repository"
	"gin-starter/modules/notification/v1/service"
)

// BuildNotificationHandler build user handlers
// starting from handler down to repository or tool.
func BuildNotificationHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool) {
	// Repository
	notificationRp := repository.NewNotificationRepository(db)

	nf := service.NewNotificationFinder(
		cfg,
		notificationRp,
	)
	nu := service.NewNotificationUpdater(
		cfg,
		notificationRp,
	)
	nc := service.NewNotificationCreator(
		cfg,
		notificationRp,
	)

	app.NotificationFinderHTTPHandler(cfg, router, nf, nu)
	app.NotificationCreatorHTTPHandler(cfg, router, nc)
	app.NotificationUpdaterHTTPHandler(cfg, router, nu)
}

// BuildSendEmailPubsubHandler is used to build the pubsub handler.
func BuildSendEmailPubsubHandler(config config.Config, db *gorm.DB) *handler.SendEmailPubsubHandler {
	repo := repository.NewEmailSent(db)
	svc := service.NewEmailSender(repo, config.MailGun)

	return handler.NewSendEmailPubsubHandler(svc, config.MailGun)
}
