package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/axelsccp/tcd-apis-amazon/src/database"

	"github.com/gorilla/mux"
)

// "Person type" (tipo um objeto)
type Item struct {
	ID    string `json:"id,omitempty"`
	Nome  string `json:"Nome,omitempty"`
	Marca string `json:"Marca,omitempty"`
	Valor string `json:"Valor,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var produto []Item

// GetPeople mostra todos os contatos da variável people
func GetProduto(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(produto)
}

// GetPerson mostra apenas um contato
func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range produto {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

// CreatePerson cria um novo contato
func CriaItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = params["id"]
	produto = append(produto, item)
	json.NewEncoder(w).Encode(produto)
}

// DeletePerson deleta um contato
func DeletaItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range produto {
		if item.ID == params["id"] {
			produto = append(produto[:index], produto[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(produto)
	}
}

// função principal para executar a api
func main() {
	var err error
	result := database.Connection("SELECT 1")
	fmt.Println(result)
	if err != nil {
		fmt.Println("error")
	}
	router := mux.NewRouter()
	produto = append(produto, Item{ID: "1", Nome: "John", Marca: "Doe", Valor: "15,00"})
	produto = append(produto, Item{ID: "2", Nome: "Koko", Marca: "Doe", Valor: "20,00"})
	router.HandleFunc("/produto", GetProduto).Methods("GET")
	router.HandleFunc("/item/{id}", GetItem).Methods("GET")
	router.HandleFunc("/item/{id}", CriaItem).Methods("POST")
	router.HandleFunc("/item/{id}", DeletaItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
