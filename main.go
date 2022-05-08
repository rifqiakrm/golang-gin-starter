package main

import (
	"context"
	"fmt"
	authBuilder "gin-starter/modules/auth/v1/builder"
	cmsBuilder "gin-starter/modules/cms/v1/builder"
	faqBuilder "gin-starter/modules/faq/v1/builder"
	notificationBuilder "gin-starter/modules/notification/v1/builder"
	pageBuilder "gin-starter/modules/page/v1/builder"
	userBuilder "gin-starter/modules/user/v1/builder"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
	pubsubSDK "gin-starter/sdk/pubsub"
	"gin-starter/utils"
)

// splash print plain text message to console
func splash() {
	fmt.Print(`
        .__                   __                 __                
   ____ |__| ____     _______/  |______ ________/  |_  ___________ 
  / ___\|  |/    \   /  ___/\   __\__  \\_  __ \   __\/ __ \_  __ \
 / /_/  >  |   |  \  \___ \  |  |  / __ \|  | \/|  | \  ___/|  | \/
 \___  /|__|___|  / /____  > |__| (____  /__|   |__|  \___  >__|   
/_____/         \/       \/            \/                 \/       
`)
}

func main() {
	cfg, err := config.LoadConfig(".env")
	checkError(err)

	splash()

	db, err := utils.NewPostgresGormDB(&cfg.Postgres)
	checkError(err)

	redisPool := buildRedisPool(cfg)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/public", "./public")

	BuildHandler(*cfg, router, db, redisPool)

	// Uncomment if you use Google pub/sub
	// psClient := createPubSubClient(cfg.Google.ProjectID, cfg.Google.ServiceAccountFile)
	// psHandlers := registerPubsubHandlers(context.Background(), cfg, db, redisPool)
	//
	// _ = psClient.StartSubscriptions(psHandlers...)

	if err := router.Run(fmt.Sprintf(":%s", cfg.Port.APP)); err != nil {
		panic(err)
	}
}

func BuildHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool) {
	app.DefaultHTTPHandler(cfg, router)
	authBuilder.BuildAuthHandler(cfg, router, db, redisPool)
	userBuilder.BuildUserHandler(cfg, router, db, redisPool)
	cmsBuilder.BuildCMSHandler(cfg, router, db, redisPool)
	notificationBuilder.BuildNotificationHandler(cfg, router, db, redisPool)
	faqBuilder.BuildFaqHandler(cfg, router, db, redisPool)
	pageBuilder.BuildPageHandler(cfg, router, db, redisPool)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// CORSMiddleware ..
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}

		c.Next()
	}
}

func registerPubsubHandlers(
	ctx context.Context,
	cfg *config.Config,
	gconn *gorm.DB,
	redisPool *redis.Pool,
) []pubsubSDK.Subscriber {
	var handlers []pubsubSDK.Subscriber

	handlers = append(handlers, notificationBuilder.BuildSendEmailPubsubHandler(*cfg, gconn))
	return handlers
}

func buildRedisPool(cfg *config.Config) *redis.Pool {
	cachePool := utils.NewPool(cfg.Redis.Address, cfg.Redis.Password)

	ctx := context.Background()
	_, err := cachePool.GetContext(ctx)

	if err != nil {
		checkError(err)
	}

	log.Print("redis successfully connected!")
	return cachePool
}

func createPubSubClient(projectID, googleSaFile string) *pubsubSDK.PubSub {
	return pubsubSDK.NewPubSub(projectID, &googleSaFile)
}
