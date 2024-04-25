package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtoHandler)
	http.ListenAndServe(":8080", nil)

}

func produtoHandler(w http.ResponseWriter, r *http.Request) {
	produto := criaEstoque()

	json.NewEncoder(w).Encode(produto)
}
