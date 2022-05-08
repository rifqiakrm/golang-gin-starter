package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/config"
	"gin-starter/middleware"
	authhandlerv1 "gin-starter/modules/auth/v1/handler"
	authservicev1 "gin-starter/modules/auth/v1/service"
	"gin-starter/modules/cms/v1/handler"
	"gin-starter/modules/cms/v1/service"
	faqhandlerv1 "gin-starter/modules/faq/v1/handler"
	faqservicev1 "gin-starter/modules/faq/v1/service"
	notificationhandlerv1 "gin-starter/modules/notification/v1/handler"
	notificationservicev1 "gin-starter/modules/notification/v1/service"
	pagehandlerv1 "gin-starter/modules/page/v1/handler"
	pageservicev1 "gin-starter/modules/page/v1/service"
	userhandlerv1 "gin-starter/modules/user/v1/handler"
	userservicev1 "gin-starter/modules/user/v1/service"
	"gin-starter/response"
)

func DeprecatedApi(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.ErrorApiResponse(http.StatusForbidden, "this version of api is deprecated. please use another version."))
	c.Abort()
	return
}

func DefaultHTTPHandler(cfg config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, response.ErrorApiResponse(http.StatusNotFound, "invalid route"))
		c.Abort()
		return
	})
}

func AuthHTTPHandler(cfg config.Config, router *gin.Engine, auc authservicev1.AuthUseCase) {
	hnd := authhandlerv1.NewAuthHandler(auc)
	v1 := router.Group("/v1")
	{
		v1.POST("/user/login", hnd.Login)
		v1.POST("/cms/login", hnd.LoginCMS)
	}
}

func UserFinderHTTPHandler(cfg config.Config, router *gin.Engine, uuc userservicev1.UserFinderUseCase) {
	hnd := userhandlerv1.NewUserFinderHandler(uuc)
	v1 := router.Group("/v1")
	{
		v1.GET("/user/forgot-password/profile/:token", hnd.GetUserByForgotPasswordToken)
	}

	v1.Use(middleware.Auth(cfg))
	{
		v1.GET("/user/profile", hnd.GetUserProfile)
	}
}

func UserUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, uuc userservicev1.UserUpdaterUseCase, uuf userservicev1.UserFinderUseCase) {
	hnd := userhandlerv1.NewUserUpdaterHandler(uuc, uuf)
	v1 := router.Group("/v1")
	{
		v1.PUT("/user/forgot-password/request", hnd.ForgotPasswordRequest)
		v1.PUT("/user/forgot-password", hnd.ForgotPassword)
	}

	v1.Use(middleware.Auth(cfg))
	{
		v1.PUT("/user/profile", hnd.UpdateUser)
		v1.PUT("/user/password", hnd.ChangePassword)
		v1.PUT("/verify/otp", hnd.VerifyOTP)
		v1.PUT("/resend/otp", hnd.ResendOTP)
	}
}

func NotificationFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationFinderUseCase, nu notificationservicev1.NotificationUpdaterUseCase) {
	hnd := notificationhandlerv1.NewNotificationFinderHandler(cf, nu)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	{
		v1.GET("/user/notifications", hnd.GetNotification)
		v1.GET("/user/notification/count", hnd.CountUnreadNotifications)
	}
}

func NotificationCreatorHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationCreatorUseCase) {
	hnd := notificationhandlerv1.NewNotificationCreatorHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.POST("/cms/notification", hnd.CreateNotification)
	}

}

func NotificationUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationUpdaterUseCase) {
	hnd := notificationhandlerv1.NewNotificationUpdaterHandler(cf)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	{
		v1.PUT("/user/notification/set", hnd.RegisterUnregisterPlayerID)
		v1.PUT("/user/notification/read", hnd.UpdateReadNotification)
	}
}

func FaqFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf faqservicev1.FaqFinderUseCase) {
	hnd := faqhandlerv1.NewFaqFinderHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.GET("/faq", hnd.GetFaq)
	}
}

func PageFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf pageservicev1.PageFinderUseCase) {
	hnd := pagehandlerv1.NewPageFinderHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.GET("/pages", hnd.GetPages)
	}
}

func CMSFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf service.CMSFinderUseCase) {
	hnd := handler.NewCMSFinderHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.GET("/cms/pages", hnd.GetPages)
		v1.GET("/cms/page/:id", hnd.GetPageByID)
	}

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.GET("/cms/profile", hnd.GetAdminProfile)
		v1.GET("/cms/admin/list", hnd.GetAdminUsers)
		v1.GET("/cms/admin/detail/:id", hnd.GetAdminUserByID)
		v1.GET("/cms/user/list", hnd.GetUsers)
		v1.GET("/cms/user/detail/:id", hnd.GetUserByID)
		v1.GET("/cms/roles", hnd.GetRoles)
		v1.GET("/cms/faqs", hnd.GetFaqs)
		v1.GET("/cms/faq/:id", hnd.GetFaqByID)
	}
}

func CMSCreatorHTTPHandler(cfg config.Config, router *gin.Engine, cc service.CMSCreatorUseCase, cf service.CMSFinderUseCase) {
	hnd := handler.NewCMSCreatorHandler(cfg, cc, cf)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.POST("/cms/user", hnd.CreateUser)
		v1.POST("/cms/admin/user", hnd.CreateAdmin)
		v1.POST("/cms/faq", hnd.CreateFaq)
		v1.POST("/cms/page", hnd.CreatePage)
	}
}

func CMSUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, cu service.CMSUpdaterUseCase, cf service.CMSFinderUseCase) {
	hnd := handler.NewCMSUpdaterHandler(cu, cf)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/cms/admin/:id", hnd.UpdateAdmin)
		v1.PUT("/cms/user/activate/:id", hnd.ActivateDeactivateUser)
		v1.PUT("/cms/faq/:id", hnd.UpdateFaq)
		v1.PUT("/cms/page/:id", hnd.UpdatePage)
	}
}

func CMSDeleterHTTPHandler(cfg config.Config, router *gin.Engine, cd service.CMSDeleterUseCase) {
	hnd := handler.NewCMSDeleterHandler(cd)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.DELETE("/cms/admin/:id", hnd.DeleteAdmin)
		v1.DELETE("/cms/faq/:id", hnd.DeleteFaq)
		v1.DELETE("/cms/page/:id", hnd.DeletePage)
	}
}
