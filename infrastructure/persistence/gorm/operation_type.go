package gorm

type OperationType struct {
	ID            int64  `gorm:"column:id;primaryKey"`
	DescriptionPT string `gorm:"column:description_pt"`
}
