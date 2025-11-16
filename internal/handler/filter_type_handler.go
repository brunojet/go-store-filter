package handler

import (
	"net/http"
	"strconv"

	"go-store-filter/internal/dto"
	"go-store-filter/internal/usecase"

	"github.com/gin-gonic/gin"
)

type FilterTypeHandler struct {
	Service *usecase.FilterTypeService
}

func NewFilterTypeHandler(service *usecase.FilterTypeService) *FilterTypeHandler {
	return &FilterTypeHandler{Service: service}
}

func (h *FilterTypeHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/tipos-filtro", h.Create)
	r.GET("/tipos-filtro", h.GetAll)
	r.GET("/tipos-filtro/:id", h.GetById)
	r.PUT("/tipos-filtro/:id", h.Update)
	r.DELETE("/tipos-filtro/:id", h.Delete)
	r.POST("/tipos-filtro/:id/tipos-filtros", h.CreateChildFilterType)
	r.GET("/tipos-filtro/:id/tipos-filtros", h.GetChildFilterTypes)
	r.POST("/tipos-filtro/:id/filtros", h.CreateFilterByFilterType)
	r.GET("/tipos-filtro/:id/filtros", h.GetFiltersByFilterType)
}

func (h *FilterTypeHandler) Create(c *gin.Context) {
	var req dto.FilterTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Conversão para model.FilterType omitida para foco no handler
	c.JSON(http.StatusCreated, gin.H{"message": "created (mock)"})
}

func (h *FilterTypeHandler) GetAll(c *gin.Context) {
	types, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, types)
}

func (h *FilterTypeHandler) GetById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	typeObj, err := h.Service.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, typeObj)
}

func (h *FilterTypeHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.FilterTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Conversão e chamada de update omitidas
	c.JSON(http.StatusOK, gin.H{"message": "updated (mock)", "id": id})
}

func (h *FilterTypeHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	// Chamada de delete omitida
	c.JSON(http.StatusOK, gin.H{"message": "deleted (mock)", "id": id})
}

// GET /tipos-filtro/:id/filtros
func (h *FilterTypeHandler) GetFiltersByFilterType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	// Aqui você buscaria os filtros pelo tipo_filtro_id
	// Exemplo mock:
	c.JSON(http.StatusOK, gin.H{"mensagem": "listar filtros do tipo_filtro", "tipo_filtro_id": id})
}

func (h *FilterTypeHandler) CreateFilterByFilterType(c *gin.Context) {
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
	// Aqui você criaria um filtro associado ao tipo_filtro_id = id
	c.JSON(http.StatusCreated, gin.H{"mensagem": "filtro criado (mock)", "tipo_filtro_id": id})
}

// GET /tipos-filtro/:id/tipos-filtros
func (h *FilterTypeHandler) GetChildFilterTypes(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	// Aqui você buscaria os tipos-filtro filhos pelo parent_id
	// Exemplo mock:
	c.JSON(http.StatusOK, gin.H{"mensagem": "listar tipos-filtros filhos", "parent_id": id})
}

// POST /tipos-filtro/:id/tipos-filtros
func (h *FilterTypeHandler) CreateChildFilterType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	var req dto.FilterTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	// Aqui você criaria um tipo-filtro filho com parent_id = id
	c.JSON(http.StatusCreated, gin.H{"mensagem": "tipo-filtro filho criado (mock)", "parent_id": id})
}

// PUT /tipos-filtro/:id/tipos-filtros/:filhoId
func (h *FilterTypeHandler) UpdateChildFilterType(c *gin.Context) {
	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	filhoId, err := strconv.ParseInt(c.Param("filhoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "filhoId inválido"})
		return
	}
	var req dto.FilterTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	// Aqui você atualizaria o tipo-filtro filho
	c.JSON(http.StatusOK, gin.H{"mensagem": "tipo-filtro filho atualizado (mock)", "filho_id": filhoId})
}

// DELETE /tipos-filtro/:id/tipos-filtros/:filhoId
func (h *FilterTypeHandler) DeleteChildFilterType(c *gin.Context) {
	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "id inválido"})
		return
	}
	filhoId, err := strconv.ParseInt(c.Param("filhoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "filhoId inválido"})
		return
	}
	// Aqui você deletaria o tipo-filtro filho
	c.JSON(http.StatusOK, gin.H{"mensagem": "tipo-filtro filho deletado (mock)", "filho_id": filhoId})
}
