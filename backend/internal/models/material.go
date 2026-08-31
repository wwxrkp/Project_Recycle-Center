package entity

type MaterialType struct {
	TypeID    int        `gorm:"column:type_id;type:integer;primaryKey;autoIncrement" json:"type_id"`
	TypeName  string     `gorm:"column:type_name;type:text;not null;uniqueIndex" json:"type_name"`
	Materials []Material `gorm:"foreignKey:MaterialTypeID;references:TypeID" json:"materials,omitempty"`
}

func (MaterialType) TableName() string { return "material_types" }

// Material is the only material master shared by Filme and Mark.
// Actual grade and condition belong to quality assessments and stock lots.
type Material struct {
	MaterialID     string        `gorm:"column:material_id;type:text;primaryKey" json:"material_id"`
	MaterialName   string        `gorm:"column:material_name;type:text;not null;uniqueIndex" json:"material_name"`
	Unit           string        `gorm:"column:unit;type:text;not null" json:"unit"`
	Status         string        `gorm:"column:status;type:text;not null" json:"status"`
	MinStockLevel  float64       `gorm:"column:min_stock_level;type:numeric(14,3);not null;default:0" json:"min_stock_level"`
	MaterialTypeID int           `gorm:"column:material_type_id;type:integer;not null;index" json:"material_type_id"`
	MaterialType   *MaterialType `gorm:"foreignKey:MaterialTypeID;references:TypeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material_type,omitempty"`
	StorageZones   []StorageZone `gorm:"foreignKey:MaterialID;references:MaterialID" json:"storage_zones,omitempty"`
}

func (Material) TableName() string { return "materials" }
