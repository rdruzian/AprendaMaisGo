package Professor

import (
	"aprendamais/Discplinas"
	"gorm.io/gorm"
	"time"
)

type Professor struct {
	ID          uint                    `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Disciplinas []Discplinas.Disciplina `json:"disciplinas" gorm:"embedded"`

	CreateAt  time.Time      `json:"created"`
	UpdateAt  time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
