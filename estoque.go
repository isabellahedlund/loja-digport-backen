package main

import "github.com/isabellahedlund/loja-digport-backend/model"

func criaEstoque() []model.Produto {
	produto := []model.Produto{
		{
			Nome:       "Disco de Vinil",
			Descricao:  "Disco de Vinil: The Beatles",
			MusicGenre: "POP",
			Artist:     "The Beatles",
			Album:      "Abbey Road",
			ID:         "1",
			Valor:      200,
			Quantidade: 1,
			Imagem:     "AbbeyRoad.png",
		},
		{
			Nome:       "Disco de Vinil",
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
			Nome:       "Disco de Vinil",
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

//criar função no catalogo que busque o produto pelo nome, chamar na main.
func buscaProdutosPorNome(nome string) []model.Produto {
	var produtosEncontrados []model.Produto

	estoque := criaEstoque()

	for _, produto := range estoque {
		if produto.Nome == nome {
			produtosEncontrados = append(produtosEncontrados, produto)
		}
	}

	return produtosEncontrados
}
