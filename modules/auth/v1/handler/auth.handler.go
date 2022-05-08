package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/modules/auth/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
)

type AuthHandler struct {
	authUseCase service.AuthUseCase
}

func NewAuthHandler(
	authUseCase service.AuthUseCase,
) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var request resource.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res, err := ah.authUseCase.AuthValidate(c, request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	token, err := ah.authUseCase.GenerateAccessToken(c, res)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	otpIsNull := false

	if res.OTP.String != "" {
		otpIsNull = true
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewLoginResponse(token.Token, otpIsNull)))
}

func (ah *AuthHandler) LoginCMS(c *gin.Context) {
	var request resource.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res, err := ah.authUseCase.AuthValidateCMS(c, request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	token, err := ah.authUseCase.GenerateAccessTokenCMS(c, res)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewLoginResponse(token.Token, false)))
}
