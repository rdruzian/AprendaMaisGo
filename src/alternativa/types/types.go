package types

type Alternativa struct {
	AlternativaTexto bool
	AlternativaImagem bool
	Texto AlternativaTexto
	Figura AlternativaImagem
	Correta string `db:resp_correta`
}

type AlternativaImagem struct {
	A []byte `db:a`
	E []byte `db:e`
	B []byte `db:b`
	C []byte `db:d`
	D []byte `db:d`
}

type AlternativaTexto struct {
	A string `db:a`
	B string `db:b`
	C string `db:d`
	D string `db:d`
	E string `db:e`
}