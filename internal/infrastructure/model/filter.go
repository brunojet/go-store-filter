package model

type Filter struct {
	Id           int64   `gorm:"primaryKey;autoIncrement;column:cod_flo" json:"-"`
	PublicId     *string `gorm:"size:3;column:cod_flo_cto" json:"filterId,omitempty"`
	Name         string  `gorm:"size:255;not null;column:nom_flo" json:"nome"`
	Description  string  `gorm:"type:text" json:"txt_desc_flo"`
	FilterTypeId int64   `gorm:"not null;column:cod_tip_flo" json:"tipoFiltroId"`
	ParentId     *int64  `gorm:"column:cod_flo_pai_id" json:"-"`
	// Foreign keys
	Parent     *Filter    `gorm:"foreignKey:ParentId" json:"-"`
	Children   []Filter   `gorm:"foreignKey:ParentId" json:"-"`
	FilterType FilterType `gorm:"foreignKey:FilterTypeId"`
}

// TableName sets the insert table name for this struct type
func (Filter) TableName() string {
	return "filtro"
}
