package dto

type FilterTypeRequest struct {
	FilterTypeId *string `json:"tipoFiltroId,omitempty"`
	Name         string  `json:"nome" validate:"required,min=1,max=255"`
	Description  string  `json:"descricao,omitempty"`
}

type FilterTypeResponse struct {
	FilterTypeId *string `json:"tipoFiltroId,omitempty"`
	Name         string  `json:"nome"`
	Description  string  `json:"descricao,omitempty"`
	ParentId     *int64  `json:"tipoFiltroPaiId,omitempty"`
}
