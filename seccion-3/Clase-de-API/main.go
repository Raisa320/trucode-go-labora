package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	//creacion del router
	router := mux.NewRouter()

	//define el comportamiento del servidor (routes)
	router.HandleFunc("/", handleRootResourcefunc).Methods("GET")

	router.HandleFunc("/otro", handleOtherResource).Methods("GET")

	router.HandleFunc("/chatgpt", handleChatGpt)
	//levantar puerto
	portNumber := 9000
	fmt.Printf("Listen in port %d", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}
}

func handleRootResourcefunc(response http.ResponseWriter, request *http.Request) {

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Soy el recurso raiz"))

}
func handleOtherResource(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in f", r)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("Respuesta fallida"))
		}
	}()

	var item *Item = nil
	fmt.Println(item.ID)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Soy otro recurso"))
}

func handleChatGpt(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("Endpoiter, request from ip %s\n", request.RemoteAddr)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()

	response.WriteHeader(http.StatusOK)
	_, err := response.Write([]byte(inputStr))
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
	}
}
