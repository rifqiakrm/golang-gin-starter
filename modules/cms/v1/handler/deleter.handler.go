package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-starter/modules/cms/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
)

type CMSDeleterHandler struct {
	cmsDeleter service.CMSDeleterUseCase
}

func NewCMSDeleterHandler(
	cmsDeleter service.CMSDeleterUseCase,
) *CMSDeleterHandler {
	return &CMSDeleterHandler{
		cmsDeleter: cmsDeleter,
	}
}

func (cc *CMSDeleterHandler) DeleteAdmin(c *gin.Context) {
	var request resource.DeleteAdminRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if err := cc.cmsDeleter.DeleteAdmin(c, uuid.MustParse(request.ID)); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (cc *CMSDeleterHandler) DeletePage(c *gin.Context) {
	var request resource.DeletePageRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	if err := cc.cmsDeleter.DeletePage(c, uuid.MustParse(request.ID)); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
	return
}

func (cc *CMSDeleterHandler) DeleteFaq(c *gin.Context) {
	var request resource.DeleteFaqRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	if err := cc.cmsDeleter.DeleteFaq(c, uuid.MustParse(request.ID)); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
	return
}
