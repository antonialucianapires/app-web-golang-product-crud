package main

import (
	"net/http"
	"app-web-golang-product-crud/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
