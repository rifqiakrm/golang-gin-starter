package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/modules/faq/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
)

type FaqFinderHandler struct {
	FaqFinder service.FaqFinderUseCase
}

func NewFaqFinderHandler(
	faqFinder service.FaqFinderUseCase,
) *FaqFinderHandler {
	return &FaqFinderHandler{
		FaqFinder: faqFinder,
	}
}

func (cf *FaqFinderHandler) GetFaq(c *gin.Context) {
	faq, total, err := cf.FaqFinder.GetFaq(c)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
		return
	}

	res := make([]*resource.Faq, 0)

	for _, n := range faq {
		res = append(res, resource.NewFaqResponse(n))
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", &resource.GetFaqResponse{
		List:  res,
		Total: total,
	}))
	return
}
