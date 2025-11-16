package dto

type FilterTypeRequest struct {
	PublicId    *string `json:"filterTypeId,omitempty"`
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Description string  `json:"description,omitempty"`
	ParentId    *int64  `json:"parentId,omitempty"`
}

type FilterTypeResponse struct {
	FilterTypeId *string `json:"filterTypeId,omitempty"`
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	ParentId     *int64  `json:"parentId,omitempty"`
}
