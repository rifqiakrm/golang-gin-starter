package handler

import (
	"encoding/base64"
	"fmt"
	service2 "gin-starter/modules/cms/v1/service"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-starter/entity"
	"gin-starter/resource"
	"gin-starter/response"
	"gin-starter/utils"
)

type CMSUpdaterHandler struct {
	cmsUpdater service2.CMSUpdaterUseCase
	cmsFinder  service2.CMSFinderUseCase
}

func NewCMSUpdaterHandler(
	cmsUpdater service2.CMSUpdaterUseCase,
	cmsFinder service2.CMSFinderUseCase,
) *CMSUpdaterHandler {
	return &CMSUpdaterHandler{
		cmsUpdater: cmsUpdater,
		cmsFinder:  cmsFinder,
	}
}

func (cc *CMSUpdaterHandler) ActivateDeactivateUser(c *gin.Context) {
	var request resource.DeactivateUserRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if err := cc.cmsUpdater.ActivateDeactivateUser(c, uuid.MustParse(request.ID)); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (cc *CMSUpdaterHandler) UpdateAdmin(c *gin.Context) {
	var request resource.UpdateAdminRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	userID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	_, err = cc.cmsFinder.GetUserByID(c, userID)

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
		userID,
		request.Name,
		request.Email,
		request.Name,
		utils.TimeToNullTime(dob),
		photo,
		request.PhoneNumber,
		"system",
	)

	roleId, err := uuid.Parse(request.RoleID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if err := cc.cmsUpdater.UpdateAdmin(c, user, roleId); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (cc *CMSUpdaterHandler) UpdateFaq(c *gin.Context) {
	var request resource.UpdateFaqRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	faqID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	faq := entity.NewFaq(
		faqID,
		request.Question,
		request.Answer,
		"system",
	)

	if err := cc.cmsUpdater.UpdateFaq(c, faq); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
	return
}

func (cc *CMSUpdaterHandler) UpdatePage(c *gin.Context) {
	var request resource.UpdatePageRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	newsCategoryID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	newsCategory := entity.NewPage(
		newsCategoryID,
		request.Title,
		request.Content,
		"system",
	)

	if err := cc.cmsUpdater.UpdatePage(c, newsCategory); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
	return
}
