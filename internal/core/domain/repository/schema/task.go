package schema

type Task struct {
	Title       string `gorm:"column:title;primaryKey" json:"title,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	Status      bool   `gorm:"column:status" json:"status,omitempty"`

	StatusData *Status `gorm:"foreignKey:Name;references:Status" json:"statusData,omitempty"`
}

func (Task) TableName() string {
	return "Task"
}
