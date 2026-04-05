package handler

import (
	"net/http"
	"search-engine/model"
	"search-engine/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// SearchHandler 搜索 HTTP 处理器
type SearchHandler struct {
	svc *service.SearchService
}

func NewSearchHandler(svc *service.SearchService) *SearchHandler {
	return &SearchHandler{svc: svc}
}

// Search GET /api/v1/search?q=关键词&page=1&size=10
func (h *SearchHandler) Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q"))
	if query == "" {
		c.JSON(http.StatusBadRequest, model.Fail("搜索关键词不能为空"))
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	resp, err := h.svc.Search(query, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.OK(resp))
}
