package routes

import 	(
	"net/http"

	"app-web-golang-product-crud/controller"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/save", controller.Insert)
}