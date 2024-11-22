package controller

import (
	"html/template"
	"net/http"

	"app-web-golang-product-crud/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscarProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		model.InserirProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	model.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    produto := model.BuscarProdutoPorID(id)
    temp.ExecuteTemplate(w, "edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        id := r.FormValue("id")
        nome := r.FormValue("nome")
        descricao := r.FormValue("descricao")
        preco := r.FormValue("preco")
        quantidade := r.FormValue("quantidade")
        dataCriacao := r.FormValue("data_criacao") // Supondo que você também esteja enviando a data de criação

        model.AtualizaProduto(id, nome, descricao, preco, quantidade, dataCriacao)
    }
    http.Redirect(w, r, "/", 301)
}