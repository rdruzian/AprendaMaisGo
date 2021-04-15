package aluno

type aluno struct {
	Nome string `db:'nome'`
	Nascimento string `db:'nascimento'`
	Area string `db:'area'`
}
