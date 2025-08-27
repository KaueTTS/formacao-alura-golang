package controllers

import (
	"encoding/json"
	"fmt"
	"go_desenvolvendo_uma_api_rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Personalidades)
}

func PersonalidadePorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	paramID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, personalidade := range models.Personalidades {
		if personalidade.Id == paramID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(personalidade)
			return
		}
	}

	http.Error(w, "Personalidade não encontrada", http.StatusNotFound)
}
