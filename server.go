package main

import (
	"encoding/json"
	"net/http"

	"github.com/isabellahedlund/loja-digport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtoHandler)
	http.ListenAndServe(":8080", nil)

}

func produtoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProduto(w, r)
	}
}

func addProduto(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	adicionaProduto(produto)

	w.WriteHeader(http.StatusCreated)
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		produtosFiltradosPorNome := buscaProdutosPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosFiltradosPorNome)
	} else {
		produtos := Produtos
		json.NewEncoder(w).Encode(produtos)
	}
}
