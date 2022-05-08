package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/modules/page/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
)

type PageFinderHandler struct {
	PageFinder service.PageFinderUseCase
}

func NewPageFinderHandler(
	pageFinder service.PageFinderUseCase,
) *PageFinderHandler {
	return &PageFinderHandler{
		PageFinder: pageFinder,
	}
}

func (cf *PageFinderHandler) GetPages(c *gin.Context) {
	pages, total, err := cf.PageFinder.GetPages(c)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	res := make([]*resource.Page, 0)

	for _, n := range pages {
		res = append(res, resource.NewPageResponse(n))
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", &resource.GetPageResponse{
		List:  res,
		Total: total,
	}))
	return
}
