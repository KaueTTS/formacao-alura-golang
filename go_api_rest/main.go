package main

import (
	"go_api_rest/database"
	"go_api_rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
