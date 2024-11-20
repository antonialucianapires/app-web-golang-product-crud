package model

import (
	"time"

	"app-web-golang-product-crud/database"
)

type Produto struct {
	Id          int
	Nome        string
	Descricao   string
	Preco       float64
	Quantidade  int
	DataCriacao time.Time
}

func BuscarProdutos() []Produto { 
	db := database.ConectaDB()
	defer db.Close() 

	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	defer selectProdutos.Close()

	produtos := []Produto{}
	for selectProdutos.Next() {
		var p Produto
		err = selectProdutos.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade, &p.DataCriacao)
		if err != nil {
			panic(err.Error())
		}
		produtos = append(produtos, p)
	}
	return produtos
}
