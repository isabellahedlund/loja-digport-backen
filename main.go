package main

import "fmt"

func main() {
	//	StartServer()
	nome := ""
	fmt.Scan("Digite o nome do produto desejado: ", &nome)

	produtosFiltrados := buscaProdutosPorNome(nome)

	fmt.Println(produtosFiltrados)
}
