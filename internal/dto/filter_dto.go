package dto

type FilterRequest struct {
	PublicId     *string `json:"filterId,omitempty"`
	Name         string  `json:"name" validate:"required,min=1,max=255"`
	FilterTypeId int64   `json:"filterTypeId" validate:"required"`
	ParentId     *int64  `json:"parentId,omitempty"`
	Value        string  `json:"value,omitempty"`
}

type FilterResponse struct {
	FilterId     *string `json:"filterId,omitempty"`
	Name         string  `json:"name"`
	FilterTypeId int64   `json:"filterTypeId"`
	ParentId     *int64  `json:"parentId,omitempty"`
	Value        string  `json:"value,omitempty"`
}
