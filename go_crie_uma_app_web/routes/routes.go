package routes

import (
	"go_crie_uma_app_web/controllers"
	"net/http"
)

func CarregaRotas() {
	// Rota da index
	http.HandleFunc("/", controllers.Index)
	// Rota da nova página
	http.HandleFunc("/new", controllers.New)
	// Rota de inserção de produto
	http.HandleFunc("/insertProduto", controllers.Insert)
	// Rota de deleção de produto
	http.HandleFunc("/deleteProduto", controllers.Delete)
	// Rota de edição de produto
	http.HandleFunc("/editProduto", controllers.Edit)
	// Rota de atualização de produto
	http.HandleFunc("/updateProduto", controllers.Update)
}
