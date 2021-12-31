package User

import (
	Areas "../Areas/areas.go"
	Disciplinas "../Areas/disciplinas.go"
)

type Usuario struct {
	ID int
	Nome string
	Cidade string
	UF string
	AreaInteresse Areas
	Tipo string
}

type Aluno struct {
	IDAluno int
	EM int
}

type Professor struct {
	IDProf int
	Disciplinas Disciplinas
}
