package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
	"time"
)

func conectaBD() *sql.DB {
	conexao := "user=user dbname=gestao_produto password=user host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id          int
	Nome        string
	Descricao   string
	Preco       float64
	Quantidade  int
	DataCriacao time.Time
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaBD()
	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectProdutos.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int
		var dataCriacao time.Time
		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade, &dataCriacao)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		p.DataCriacao = dataCriacao
		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}
