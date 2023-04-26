package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	router.HandleFunc("/item", searchItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	portNumber := 9000
	fmt.Printf("Listen in port %d\n", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}
}

func getItems(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	// Función para obtener todos los elementos
	json.NewEncoder(response).Encode(items)
}

func searchItem(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	paramName := request.URL.Query()
	name, err := paramName["name"]
	if !err {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	var itemsFound []Item
	for _, item := range items {
		if strings.Contains(item.Name, name[0]) {
			itemsFound = append(itemsFound, item)
		}
	}
	json.NewEncoder(response).Encode(itemsFound)
}

func getItem(response http.ResponseWriter, request *http.Request) {
	// Función para obtener un elemento específico
	response.Header().Set("Content-Type", "application/json")
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

func createItem(response http.ResponseWriter, request *http.Request) {
	// Función para crear un nuevo elemento
	response.Header().Set("Content-Type", "application/json")
	var item Item
	err := json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	items = append(items, item)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(items)
}

func updateItem(response http.ResponseWriter, request *http.Request) {
	// Función para actualizar un elemento existente
	response.Header().Set("Content-Type", "application/json")
	var itemUpdate Item
	err := json.NewDecoder(request.Body).Decode(&itemUpdate)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	param := mux.Vars(request)
	idItem := param["id"]
	for indice, item := range items {
		if item.ID == idItem {
			items[indice] = itemUpdate
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(items)
			return
		}
	}
	response.WriteHeader(http.StatusNotFound)
	json.NewEncoder(response).Encode(Response{"item not found"})
}

func deleteItem(response http.ResponseWriter, request *http.Request) {
	// Función para eliminar un elemento
	response.Header().Set("Content-Type", "application/json")
	param := mux.Vars(request)
	idItem := param["id"]
	for indice, item := range items {
		if item.ID == idItem {
			items = append(items[:indice], items[indice+1:]...)
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(items)
			return
		}
	}
	response.WriteHeader(http.StatusNotFound)
	json.NewEncoder(response).Encode(Response{"item not found"})

}
