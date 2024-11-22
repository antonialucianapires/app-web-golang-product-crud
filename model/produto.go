package model

import (
	"log"
	"strconv"
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

func InserirProduto(nome, descricao, preco, quantidade string) {
    precoConvertido, err := strconv.ParseFloat(preco, 64)
    if err != nil {
        log.Println("Erro na conversão do preço:", err)
    }

    quantidadeConvertida, err := strconv.Atoi(quantidade)
    if err != nil {
        log.Println("Erro na conversão da quantidade:", err)
    }

    db := database.ConectaDB()
    defer db.Close()

    insertProduto, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade, data_criacao) values($1, $2, $3, $4, $5)")
    if err != nil {
        panic(err.Error())
    }
    defer insertProduto.Close()

    insertProduto.Exec(nome, descricao, precoConvertido, quantidadeConvertida, time.Now())
}

func DeletaProduto (id string) {
	db := database.ConectaDB()
	defer db.Close()

	deleteProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	defer deleteProduto.Close()

	deleteProduto.Exec(id)
}

func BuscarProdutoPorID(id string) Produto {
    db := database.ConectaDB()
    defer db.Close()

    var p Produto
    err := db.QueryRow("select id, nome, descricao, preco, quantidade, data_criacao from produtos where id=$1", id).Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade, &p.DataCriacao)
    if err != nil {
        panic(err.Error())
    }

    return p
}

func AtualizaProduto(id, nome, descricao, preco, quantidade, dataCriacao string) {
    precoConvertido, err := strconv.ParseFloat(preco, 64)
    if err != nil {
        log.Println("Erro na conversão do preço:", err)
    }

    quantidadeConvertida, err := strconv.Atoi(quantidade)
    if err != nil {
        log.Println("Erro na conversão da quantidade:", err)
    }

    dataCriacaoConvertida, err := time.Parse("2006-01-02 15:04:05", dataCriacao)
    if err != nil {
        log.Println("Erro na conversão da data de criação:", err)
    }

    db := database.ConectaDB()
    defer db.Close()

    atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4, data_criacao=$5 where id=$6")
    if err != nil {
        panic(err.Error())
    }
    defer atualizaProduto.Close()

    atualizaProduto.Exec(nome, descricao, precoConvertido, quantidadeConvertida, dataCriacaoConvertida, id)
}