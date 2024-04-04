package model

type Carrinho struct {
	ID           string
	UserID       string
	InfosProduto []InfosProduto
	ValorTotal   float64
}

type InfosProduto struct {
	ProdutoId           string
	QuantidadeDeProduto int
}
