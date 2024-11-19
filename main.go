package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index (w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Camiseta de algodão, ideal para o dia a dia", 49.90, 12},
		{"Calça Jeans", "Calça jeans confortável com corte reto", 89.90, 8},
		{"Tênis Casual", "Tênis casual, ideal para passeios e trabalho", 129.90, 5},
	}


	temp.ExecuteTemplate(w, "index", produtos)
}