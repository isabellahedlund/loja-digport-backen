package main

import (
	"errors"

	"github.com/isabellahedlund/loja-digport-backend/model"
)

var Produtos []model.Produto = []model.Produto{}

func criaEstoque() []model.Produto {
	produto := []model.Produto{
		{
			Nome:       "Abbey",
			Descricao:  "Disco de Vinil: The Beatles, Abbey Road",
			MusicGenre: "POP",
			Artist:     "The Beatles",
			Album:      "Abbey Road",
			ID:         "1",
			Valor:      200,
			Quantidade: 1,
			Imagem:     "AbbeyRoad.png",
		},
		{
			Nome:       "Help",
			Descricao:  "Disco de Vinil: The Beatles",
			MusicGenre: "POP",
			Artist:     "The Beatles",
			Album:      "Help",
			ID:         "2",
			Valor:      150,
			Quantidade: 1,
			Imagem:     "Help.png",
		},
		{
			Nome:       "The Beatles",
			Descricao:  "Disco de Vinil: The Beatles",
			MusicGenre: "POP",
			Artist:     "The Beatles",
			Album:      "",
			ID:         "3",
			Valor:      150,
			Quantidade: 1,
			Imagem:     "Help.png",
		},
	}
	return produto
}

//func buscaProdutosPorNome(nome string) []model.Produto

// criar função no catalogo que busque o produto pelo nome, chamar na main.
func buscaProdutosPorNome(nome string) []model.Produto {

	estoque := []model.Produto{}

	for _, produto := range criaEstoque() {
		if produto.Nome == nome {
			estoque = append(estoque, produto)
		}
	}

	return estoque
}

func adicionaProduto(produtoNovo model.Produto) error {
	for _, produto := range Produtos {
		if produtoNovo.ID == produto.ID {
			return errors.New("produto com ID já existente")
		}
	}
	Produtos = append(Produtos, produtoNovo)
	return nil
}
