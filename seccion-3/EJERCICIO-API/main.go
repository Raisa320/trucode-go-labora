package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

var items = []Item{
	{
		ID:   "1",
		Name: "Item 1",
	},
	{
		ID:   "2",
		Name: "Item 2",
	},
	{
		ID:   "3",
		Name: "Item 3",
	},
	{
		ID:   "4",
		Name: "Item 4",
	},
	{
		ID:   "5",
		Name: "Item 5",
	},
	{
		ID:   "6",
		Name: "Item 6",
	},
	{
		ID:   "7",
		Name: "Item 7",
	},
	{
		ID:   "8",
		Name: "Item 8",
	},
	{
		ID:   "9",
		Name: "Item 9",
	},
	{
		ID:   "10",
		Name: "Item 10",
	},
}

func main() {
	router := mux.NewRouter()

	//Tarea.
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")

	//Más adelante.
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	portNumber := 9000
	fmt.Printf("Listen in port %d", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}
}

func getItems(response http.ResponseWriter, request *http.Request) {
	// Función para obtener todos los elementos
	json.NewEncoder(response).Encode(items)
}

func getItem(response http.ResponseWriter, request *http.Request) {
	// Función para obtener un elemento específico
	param := mux.Vars(request)
	idItem := param["id"]
	for _, item := range items {
		if item.ID == idItem {
			json.NewEncoder(response).Encode(item)
			return
		}
	}
	response.WriteHeader(http.StatusNotFound)
	json.NewEncoder(response).Encode(Response{"item not found"})
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Función para crear un nuevo elemento
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
}
