package schema

type Status struct {
	Name        string `gorm:"column:name;primaryKey" json:"name,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	Active      bool   `gorm:"column:active" json:"active,omitempty"`
}

func (Status) TableName() string {
	return "status"
}
