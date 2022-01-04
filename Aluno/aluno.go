package Aluno

import (
	"gorm.io/gorm"
	"time"
)

type Aluno struct {
	ID uint `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	EM int

	CreateAt  time.Time      `json:"created"`
	UpdateAt  time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
