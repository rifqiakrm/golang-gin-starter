package handler

import (
	"gin-starter/modules/user/v1/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/middleware"
	"gin-starter/resource"
	"gin-starter/response"
)

type UserFinderHandler struct {
	userFinder service.UserFinderUseCase
}

func NewUserFinderHandler(
	userFinder service.UserFinderUseCase,
) *UserFinderHandler {
	return &UserFinderHandler{
		userFinder: userFinder,
	}
}

func (ah *UserFinderHandler) GetUserProfile(c *gin.Context) {
	res, err := ah.userFinder.GetUserByID(c, middleware.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewUserProfile(res)))
}

func (ah *UserFinderHandler) GetUserByForgotPasswordToken(c *gin.Context) {
	var request resource.GetUserByForgotPasswordTokenRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res, err := ah.userFinder.GetUserByForgotPasswordToken(c, request.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewUserProfile(res)))
}
