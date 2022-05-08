package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-starter/middleware"
	"gin-starter/modules/cms/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
)

type CMSFinderHandler struct {
	cmsFinder service.CMSFinderUseCase
}

func NewCMSFinderHandler(
	cmsFinder service.CMSFinderUseCase,
) *CMSFinderHandler {
	return &CMSFinderHandler{
		cmsFinder: cmsFinder,
	}
}

func (cf *CMSFinderHandler) GetAdminProfile(c *gin.Context) {
	user, err := cf.cmsFinder.GetUserByID(c, middleware.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewUserAdmin(user)))
}

func (cf *CMSFinderHandler) GetUsers(c *gin.Context) {
	var request resource.GetAdminUsersRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	users, total, err := cf.cmsFinder.GetUsers(c, request.Query, request.Sort, request.Order, request.Limit, request.Offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.UserProfile, 0)

	for _, u := range users {
		res = append(res, resource.NewUserProfile(u))
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", &resource.GetUsersResponse{
		List:  res,
		Total: total,
	}))
}

func (cf *CMSFinderHandler) GetUserByID(c *gin.Context) {
	var request resource.GetUserByIDRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	user, err := cf.cmsFinder.GetUserByID(c, uuid.MustParse(request.ID))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewUserProfile(user)))
}

func (cf *CMSFinderHandler) GetAdminUsers(c *gin.Context) {
	var request resource.GetAdminUsersRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	users, total, err := cf.cmsFinder.GetAdminUsers(c, request.Query, request.Sort, request.Order, request.Limit, request.Offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.UserAdmin, 0)

	for _, u := range users {
		res = append(res, resource.NewUserAdmin(u))
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", &resource.GetAdminUsersResponse{
		List:  res,
		Total: total,
	}))
}

func (cf *CMSFinderHandler) GetAdminUserByID(c *gin.Context) {
	var request resource.GetAdminUserByIDRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	userID, err := uuid.Parse(request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	user, err := cf.cmsFinder.GetUserByID(c, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewUserAdmin(user)))
}

func (cf *CMSFinderHandler) GetRoles(c *gin.Context) {
	page, err := cf.cmsFinder.GetRoles(c)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.Role, 0)

	for _, v := range page {
		res = append(res, resource.NewRoleResponse(v))
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", res))
}

func (cf *CMSFinderHandler) GetFaqs(c *gin.Context) {
	faq, err := cf.cmsFinder.GetFaqs(c)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	res := make([]*resource.Faq, 0)

	for _, v := range faq {
		res = append(res, resource.NewFaqResponse(v))
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", res))
	return
}

func (cf *CMSFinderHandler) GetFaqByID(c *gin.Context) {
	var request resource.GetFaqByID

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	faq, err := cf.cmsFinder.GetFaqByID(c, uuid.MustParse(request.ID))

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewFaqResponse(faq)))
	return
}

func (cf *CMSFinderHandler) GetPages(c *gin.Context) {
	page, err := cf.cmsFinder.GetPages(c)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	res := make([]*resource.Page, 0)

	for _, v := range page {
		res = append(res, resource.NewPageResponse(v))
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", res))
	return
}

func (cf *CMSFinderHandler) GetPageByID(c *gin.Context) {
	var request resource.GetPageByIDRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	page, err := cf.cmsFinder.GetPageByID(c, uuid.MustParse(request.ID))

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", resource.NewPageResponse(page)))
	return
}
