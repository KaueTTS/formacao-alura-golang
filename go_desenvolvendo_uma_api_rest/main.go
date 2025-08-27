package main

import (
	"fmt"
	"go_desenvolvendo_uma_api_rest/database"
	"go_desenvolvendo_uma_api_rest/models"
	"go_desenvolvendo_uma_api_rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
		{Id: 3, Nome: "Nome 3", Historia: "Historia 3"},
	}

	database.ConectaBanco()
	fmt.Println("Iniciando o servidor na porta 8000...")
	routes.HandleRequest()
}
