package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-starter/config"
	"gin-starter/response"
	"gin-starter/utils"
)

var UserID uuid.UUID

func Auth(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(tokenString) < 2 {
			c.JSON(http.StatusUnauthorized, response.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		claims, err := utils.JWTDecode(cfg, tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorApiResponse(http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		UserID = claims.Subject

		c.Next()
	}
}

func Admin(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(tokenString) < 2 {
			c.JSON(http.StatusUnauthorized, response.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		claims, err := utils.JWTDecode(cfg, tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorApiResponse(http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		if claims.Issuer != cfg.JWTConfig.IssuerCMS {
			c.JSON(http.StatusUnauthorized, response.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		UserID = claims.Subject

		c.Next()
	}
}
