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