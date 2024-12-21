package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializa o roteador
	router := mux.NewRouter()

	// Define uma rota de exemplo
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bem-vindo ao Portal da ONG!"))
	})

	// Inicia o servidor
	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
