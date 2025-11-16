package dto

type FilterRequest struct {
	FilterId    *string `json:"filtroId,omitempty"`
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Description string  `json:"description,omitempty"`
}

type FilterResponse struct {
	FilterId     *string `json:"filtroId,omitempty"`
	Name         string  `json:"nome"`
	Description  string  `json:"descricao,omitempty"`
	FilterTypeId int64   `json:"tipoFiltroId"`
	ParentId     *int64  `json:"filtroPaiId,omitempty"`
}
