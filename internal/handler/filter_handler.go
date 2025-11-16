package handler

import (
	"net/http"
	"strconv"

	"go-store-filter/internal/dto"
	"go-store-filter/internal/usecase"

	"github.com/gin-gonic/gin"
)

type FilterHandler struct {
	Service *usecase.FilterService
}

func NewFilterHandler(service *usecase.FilterService) *FilterHandler {
	return &FilterHandler{Service: service}
}

func (h *FilterHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/filtros", h.Create)
	r.GET("/filtros", h.GetAll)
	r.GET("/filtros/:id", h.GetById)
	r.PUT("/filtros/:id", h.Update)
	r.DELETE("/filtros/:id", h.Delete)
	r.POST("/filtros/:id/filtros", h.CreateChildFilter)
	r.GET("/filtros/:id/filtros", h.GetChildFilters)
}

func (h *FilterHandler) Create(c *gin.Context) {
	var req dto.FilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Conversão para model.Filter omitida para foco no handler
	// Adapte conforme necessário
	c.JSON(http.StatusCreated, gin.H{"message": "created (mock)"})
}

// GET /filtros
func (h *FilterHandler) GetAll(c *gin.Context) {
	filters, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, filters)
}

func (h *FilterHandler) GetById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	filter, err := h.Service.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, filter)
}

func (h *FilterHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.FilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Conversão e chamada de update omitidas
	c.JSON(http.StatusOK, gin.H{"message": "updated (mock)", "id": id})
}

func (h *FilterHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	// Chamada de delete omitida
	c.JSON(http.StatusOK, gin.H{"message": "deleted (mock)", "id": id})
}

// GET /filtros/:id/filtros
func (h *FilterHandler) GetChildFilters(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	// Aqui você buscaria os filtros filhos pelo parent_id
	// Exemplo mock:
	c.JSON(http.StatusOK, gin.H{"mensagem": "listar filtros filhos", "parent_id": id})
}

// POST /filtros/:id/filtros
func (h *FilterHandler) CreateChildFilter(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	var req dto.FilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	// Aqui você criaria um filtro filho com parent_id = id
	c.JSON(http.StatusCreated, gin.H{"mensagem": "filtro filho criado (mock)", "parent_id": id})
}
