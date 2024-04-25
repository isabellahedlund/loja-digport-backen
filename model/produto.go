package model

type Produto struct {
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	MusicGenre string  `json:"musicgenre"`
	Artist     string  `json:"artist"`
	Album      string  `json:"album"`
	ID         string  `json:"id"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
	Imagem     string  `json:"imagem"`
}
