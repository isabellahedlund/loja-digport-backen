package model

type Carrinho struct {
	UserId       string
	Quantidade   int
	ValorTotal   float64
	Id           string
	InfosProduto []InfosProduto
}

type InfosProduto struct {
	ProdutoId           string
	QuantidadeDeProduto int
}
