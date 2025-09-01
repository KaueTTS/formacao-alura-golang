package controllers

import (
	"encoding/json"
	"fmt"
	"go_desenvolvendo_uma_api_rest/database"
	"go_desenvolvendo_uma_api_rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func PersonalidadePorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	if personalidade.Id != 0 {
		json.NewEncoder(w).Encode(personalidade)
		return
	} else {
		json.NewEncoder(w).Encode("Personalidade não encontrada!")
		return
	}
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)

	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var deletaPersonalidade models.Personalidade
	database.DB.Delete(&deletaPersonalidade, id)

	json.NewEncoder(w).Encode("A personalidade foi deletada com sucesso!")
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidadeAtualizada models.Personalidade
	database.DB.First(&personalidadeAtualizada, id)
	json.NewDecoder(r.Body).Decode(&personalidadeAtualizada)
	database.DB.Save(&personalidadeAtualizada)
	json.NewEncoder(w).Encode(personalidadeAtualizada)
}
