package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Endpoint para obtener los datos de un usuario
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Simulación de datos de usuario
	user := User{
		Name:  "Juan",
		Email: "juan@example.com",
	}

	// Convertir la estructura a un objeto JSON
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}

func main() {
	// Inicializar el enrutador
	r := mux.NewRouter()

	// Definir la ruta para obtener los datos de un usuario
	r.HandleFunc("/user", getUser).Methods("GET")

	// Agregar el middleware CORS a todas las rutas
	handler := cors.Default().Handler(r)

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8000", handler))
}
