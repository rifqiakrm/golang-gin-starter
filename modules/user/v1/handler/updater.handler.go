package handler

import (
	"encoding/base64"
	"fmt"
	service2 "gin-starter/modules/user/v1/service"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gin-starter/entity"
	"gin-starter/middleware"
	"gin-starter/resource"
	"gin-starter/response"
	"gin-starter/utils"
)

type UserUpdaterHandler struct {
	userUpdater service2.UserUpdaterUseCase
	userFinder  service2.UserFinderUseCase
}

func NewUserUpdaterHandler(
	userUpdater service2.UserUpdaterUseCase,
	userFinder service2.UserFinderUseCase,
) *UserUpdaterHandler {
	return &UserUpdaterHandler{
		userUpdater: userUpdater,
		userFinder:  userFinder,
	}
}

func (ah *UserUpdaterHandler) ChangePassword(c *gin.Context) {
	var request resource.ChangePasswordRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if request.NewPassword != request.NewPasswordConfirmation {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, entity.ErrWrongPasswordConfirmation.Message))
		c.Abort()
		return
	}

	if err := ah.userUpdater.ChangePassword(
		c,
		middleware.UserID,
		request.OldPassword,
		request.NewPassword,
	); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (ah *UserUpdaterHandler) ForgotPasswordRequest(c *gin.Context) {
	var request resource.ForgotPasswordRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if err := ah.userUpdater.ForgotPasswordRequest(
		c,
		request.Email,
	); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (ah *UserUpdaterHandler) VerifyOTP(c *gin.Context) {
	var request resource.VerifyOTPRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	verify, err := ah.userUpdater.VerifyOTP(
		c,
		middleware.UserID,
		request.Code,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if !verify {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, entity.ErrOTPMismatch.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (ah *UserUpdaterHandler) ResendOTP(c *gin.Context) {
	if err := ah.userUpdater.ResendOTP(
		c,
		middleware.UserID,
	); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (ah *UserUpdaterHandler) ForgotPassword(c *gin.Context) {
	var request resource.ForgotPasswordChangeRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if request.NewPassword != request.NewPasswordConfirmation {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, entity.ErrWrongPasswordConfirmation.Message))
		c.Abort()
		return
	}

	res, err := ah.userFinder.GetUserByForgotPasswordToken(c, request.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if err := ah.userUpdater.ForgotPassword(
		c,
		res.ID,
		request.NewPassword,
	); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (uu *UserUpdaterHandler) UpdateUser(c *gin.Context) {
	var request resource.UpdateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	_, err := uu.userFinder.GetUserByID(c, middleware.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	photo := ""

	if request.Photo != nil {
		splitName := strings.Split(request.Photo.Filename, ".")
		newFilename := base64.StdEncoding.EncodeToString([]byte(splitName[0] + time.Now().Format("2006-01-02 15:04:05")))

		path := "images/photo/" + fmt.Sprintf("%s.%s", newFilename, splitName[1])
		if err := c.SaveUploadedFile(request.Photo, "./public/"+path); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
			c.Abort()
			return
		}

		photo = path
	}

	dob, err := utils.DateStringToTime(request.DOB)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	user := entity.NewUser(
		middleware.UserID,
		request.Name,
		request.Email,
		request.Name,
		utils.TimeToNullTime(dob),
		photo,
		request.PhoneNumber,
		"system",
	)

	if err := uu.userUpdater.Update(c, user); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}
