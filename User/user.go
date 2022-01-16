package User

import (
	"../Aluno"
	"../Professor"
	"gorm.io/gorm"
	"time"
)

type Usuario struct {
	ID            uint                `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Nome          string              `json:"nome"`
	Cidade        string              `json:"cidade"`
	UF            string              `json:"uf"`
	AreaInteresse int                 `json:"area-interesse"`
	Aluno         Aluno.Aluno         `json:"aluno" gorm:"foreignKey:ID" gorm:"embedded"`
	Professor     Professor.Professor `json:"professor" gorm:"foreignKey:ID" gorm:"embedded"`
	Login         string              `json:"login"`
	Senha         string              `json:"senha"`

	CreateAt  time.Time      `json:"created"`
	UpdateAt  time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
