package main

import (
	"testes1/src/api"
	postgresdb "testes1/src/config/db"
)

func main() {
	r := api.Init()

	defer postgresdb.Close()

	r.Run()
}
