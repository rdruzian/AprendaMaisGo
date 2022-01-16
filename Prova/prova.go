package Prova

import (
	"aprendamais/Alternativa/types"
	"time"
)

type Prova struct {
	Questao      QuestaoAlternativa
	Universidade Universidade
	Ano          time.Time `db:ano`
}

type QuestaoAlternativa struct {
	Pergunta        string `db:'pergunta'`
	Alternativa     types.Alternativa
	PossuiTexto     bool
	Texto           string `db:texto`
	PossuiImagem    bool
	Imagem          []byte `db:figura`
	TextoPerguntas  bool
	ImagemPerguntas bool
	Dicsciplina     Disciplina
}

type Universidade struct {
	IdUniversidade int    `db:id_universidade`
	Nome           string `db:nome`
}

type Disciplina struct {
	IdDisciplina int    `db:id_disciplina`
	Nome         string `db:nome`
}
