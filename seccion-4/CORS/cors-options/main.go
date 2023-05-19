package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Definir una estructura para los datos del usuario
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Endpoint para obtener los datos de un usuario
func getUser(w http.ResponseWriter, r *http.Request) {
	// Simulación de datos de usuario
	user := User{
		Name:  "Juan",
		Email: "juan@example.com",
	}

	// Convertir la estructura a un objeto JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Escribir la respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	// Inicializar el enrutador
	r := mux.NewRouter()

	// Definir la ruta para obtener los datos de un usuario
	r.HandleFunc("/user", getUser).Methods("GET")

	// Configurar el middleware CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://example.com"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// Agregar el middleware CORS a todas las rutas
	handler := corsOptions.Handler(r)

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8000", handler))
}
