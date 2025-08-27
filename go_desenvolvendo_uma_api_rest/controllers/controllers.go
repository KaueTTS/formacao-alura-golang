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
	w.Header().Set("Content-Type", "application/json")

	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func PersonalidadePorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}
