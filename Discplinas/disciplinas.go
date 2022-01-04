package Discplinas

import (
	"gorm.io/gorm"
	"time"
)

const (
	Biologicas = iota
	Exatas
	Humanas
)

type Disciplina struct {
	ID   uint   `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Nome string `json:"nome"`
	Area int    `json:"area"`

	CreateAt  time.Time      `json:"created"`
	UpdateAt  time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
