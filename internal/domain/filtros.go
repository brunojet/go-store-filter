package domain

type Filtros struct {
	ID           int64  `json:"id"`
	Nome         string `json:"nome"`
	TipoFiltroID int64  `json:"tipo_filtro_id"`
	Valor        string `json:"valor"`
}
