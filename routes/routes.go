package routes

import 	(
	"net/http"

	"app-web-golang-product-crud/controller"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
}