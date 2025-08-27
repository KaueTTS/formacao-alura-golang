package routes

import (
	"go_desenvolvendo_uma_api_rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.Personalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadePorID).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
