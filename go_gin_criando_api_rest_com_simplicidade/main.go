package main

import (
	"go_gin_criando_api_rest_com_simplicidade/src/api"
)

func main() {
	r := api.Init()

	r.Run(":8080")
}
