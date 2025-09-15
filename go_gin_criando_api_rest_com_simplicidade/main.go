package main

import (
	"go_gin_criando_api_rest_com_simplicidade/src/api"
	postgresdb "go_gin_criando_api_rest_com_simplicidade/src/config/db"
)

func main() {
	r := api.Init()

	defer postgresdb.Close()

	r.Run()
}
