package handler

import (
	"net/http"
	"search-engine/model"
	"search-engine/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DocHandler 文档管理 HTTP 处理器
type DocHandler struct {
	svc *service.DocService
}

func NewDocHandler(svc *service.DocService) *DocHandler {
	return &DocHandler{svc: svc}
}

// CreateDoc POST /api/v1/docs
func (h *DocHandler) CreateDoc(c *gin.Context) {
	var req model.CreateDocRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Fail("参数错误: "+err.Error()))
		return
	}

	doc, err := h.svc.CreateDoc(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.OK(doc))
}

// ListDocs GET /api/v1/docs
func (h *DocHandler) ListDocs(c *gin.Context) {
	docs, err := h.svc.GetAllDocs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OK(docs))
}

// DeleteDoc DELETE /api/v1/docs/:id
func (h *DocHandler) DeleteDoc(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Fail("无效的 ID"))
		return
	}

	if err := h.svc.DeleteDoc(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.OK(nil))
}

// RebuildIndex POST /api/v1/index/rebuild
func (h *DocHandler) RebuildIndex(c *gin.Context) {
	count, err := h.svc.RebuildIndex()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.OK(gin.H{"indexed": count}))
}
