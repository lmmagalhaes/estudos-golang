package models

import (
	"curso_api_golang/db"
)

type Produto struct {
    Id         int  
    Nome       string
    Descricao  string
    Preco      float64
    Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBanco()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
	  panic(err)
	}
	p := Produto{}
	produtos := []Produto{}
	
	for selectDeTodosOsProdutos.Next() {
	  var id, quantidade int
	  var nome, descricao string
	  var preco float64
	  
	  err := selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
	  if err != nil {
		panic(err.Error())
	  }
	 
	  p.Id = id
	  p.Nome = nome
	  p.Preco = preco
	  p.Descricao = descricao
	  p.Quantidade = quantidade
  
	  produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()
	insertStatement := `
	INSERT INTO produtos (nome, descrição, preço, quantidade)
	VALUES ($1, $2, $3, $4)`

	insereDadosNoBanco, err := db.Prepare(insertStatement)
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBanco()
	insertStatement := `
	DELETE FROM produtos WHERE id=$1`

	deletaDoBanco, err := db.Prepare(insertStatement)
	if err != nil {
		panic(err.Error())
	}
	deletaDoBanco.Exec(id)

	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBanco()
	insertStatement := `
	SELECT * FROM produtos WHERE id=$1`

	produtoDoBanco, err := db.Query(insertStatement, id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		
		err := produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
		  panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	  }
	  defer db.Close()
	  return produtoParaAtualizar
}

func AtualizaNovoProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()
	insertStatement := `
	UPDATE produtos SET nome=$1, descrição=$2, preço=$3, quantidade=$4 WHERE id=$5`

	atualizaProduto, err := db.Prepare(insertStatement)
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}