package Migrations

import (
	"../../Aluno"
	"../../Discplinas"
	"../../Professor"
	"../../User"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(User.Usuario{})
	if err != nil {
		log.Printf("Erro ao criar tabela Usuario")
		return
	}
	err = db.AutoMigrate(Aluno.Aluno{})
	if err != nil {
		log.Printf("Erro ao criar tabela Aluno")
		return
	}
	err = db.AutoMigrate(Professor.Professor{})
	if err != nil {
		log.Printf("Erro ao criar tabela Professor")
		return
	}
	err = db.AutoMigrate(Discplinas.Disciplina{})
	if err != nil {
		log.Printf("Erro ao criar tabela Disciplina")
		return
	}
}
