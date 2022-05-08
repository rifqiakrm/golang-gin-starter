package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-starter/config"
	service2 "gin-starter/modules/cms/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"gin-starter/utils"
)

type CMSCreatorHandler struct {
	cfg        config.Config
	cmsCreator service2.CMSCreatorUseCase
	cmsFinder  service2.CMSFinderUseCase
}

func NewCMSCreatorHandler(
	cfg config.Config,
	cmsCreator service2.CMSCreatorUseCase,
	cmsFinder service2.CMSFinderUseCase,
) *CMSCreatorHandler {
	return &CMSCreatorHandler{
		cfg:        cfg,
		cmsCreator: cmsCreator,
		cmsFinder:  cmsFinder,
	}
}

func (cc *CMSCreatorHandler) CreateUser(c *gin.Context) {
	var request resource.CreateUserRequest

	if err := c.ShouldBind(&request); err != nil {
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

	user, err := cc.cmsCreator.CreateUser(
		c,
		request.Name,
		request.Email,
		request.Password,
		request.PhoneNumber,
		photo,
		dob,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewUserProfile(user)))
}

func (cc *CMSCreatorHandler) CreateAdmin(c *gin.Context) {
	var request resource.CreateAdminRequest

	if err := c.ShouldBind(&request); err != nil {
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

	roleID, err := uuid.Parse(request.RoleID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	user, err := cc.cmsCreator.CreateAdmin(
		c,
		request.Name,
		request.Email,
		request.Password,
		request.PhoneNumber,
		photo,
		dob,
		roleID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewUserProfile(user)))
}

func (cc *CMSCreatorHandler) CreatePage(c *gin.Context) {
	var request resource.CreatePageRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	page, err := cc.cmsCreator.CreatePage(
		c,
		request.Title,
		request.Content,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewPageResponse(page)))
	return
}

func (cc *CMSCreatorHandler) CreateFaq(c *gin.Context) {
	var request resource.CreateFaqRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	page, err := cc.cmsCreator.CreateFaq(
		c,
		request.Question,
		request.Answer,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewFaqResponse(page)))
	return
}
