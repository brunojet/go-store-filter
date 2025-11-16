package model

type FilterType struct {
	Id          int64   `gorm:"primaryKey;autoIncrement;column:cod_tip_flo" json:"-"`
	PublicId    *string `gorm:"size:3;column:cod_tip_flo_cto" json:"tipFiltroId,omitempty"`
	Name        string  `gorm:"size:255;not null" json:"nom_flo"`
	Description string  `gorm:"type:text" json:"txt_desc_tip_flo"`
	ParentId    *int64  `gorm:"column:cod_tip_flo_pai_id" json:"-"`
	// Foreign keys
	Parent   *FilterType  `gorm:"foreignKey:ParentId" json:"-"`
	Children []FilterType `gorm:"foreignKey:ParentId" json:"-"`
	Filters  []Filter     `gorm:"foreignKey:FilterTypeId"`
}

// TableName sets the insert table name for this struct type
func (FilterType) TableName() string {
	return "tipo_filtro"
}
