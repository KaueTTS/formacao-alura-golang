package main

import (
	"fmt"
	"go_desenvolvendo_uma_api_rest/database"
	"go_desenvolvendo_uma_api_rest/routes"
)

func main() {
	database.ConectaBanco()
	fmt.Println("Iniciando o servidor na porta 8000...")
	routes.HandleRequest()
}
