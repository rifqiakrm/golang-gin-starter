package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
	authBuilder "gin-starter/modules/auth/v1/builder"
	cmsBuilder "gin-starter/modules/cms/v1/builder"
	faqBuilder "gin-starter/modules/faq/v1/builder"
	notificationBuilder "gin-starter/modules/notification/v1/builder"
	pageBuilder "gin-starter/modules/page/v1/builder"
	userBuilder "gin-starter/modules/user/v1/builder"
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
	tracer, closer, _ := NewJaegerTracer(cfg.AppName, fmt.Sprintf("%s:%s", cfg.Jaeger.Address, cfg.Jaeger.Port))

	defer func() {
		if err := closer.Close(); err != nil {
			log.Println("failed to close opentracing closer:", err)
		}
	}()

	opentracing.SetGlobalTracer(tracer)

	router.Use(OpenTracing())
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/healthz", HealthGET)

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

func NewJaegerTracer(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  jaegerHostPort,
		},

		ServiceName: serviceName,
	}

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	return tracer, closer, err
}

func OpenTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		wireCtx, _ := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header))

		serverSpan := opentracing.StartSpan(c.Request.URL.Path,
			ext.RPCServerOption(wireCtx))
		defer serverSpan.Finish()
		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), serverSpan))
		c.Next()
	}
}

func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
