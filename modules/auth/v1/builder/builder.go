package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"

	"gin-starter/app"
	"gin-starter/config"
	authRepo "gin-starter/modules/auth/v1/repository"
	auth "gin-starter/modules/auth/v1/service"
)

// BuildAuthHandler build auth handlers
// starting from handler down to repository or tool.
func BuildAuthHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool) {
	// Repository
	ar := authRepo.NewAuthRepository(db)

	uc := auth.NewAuthService(cfg, ar)

	app.AuthHTTPHandler(cfg, router, uc)
}
