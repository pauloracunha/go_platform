package main

import (
	"fmt"
	"log"
	"net/http"
	"portal/config"
	"portal/internal/routes"
	"portal/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// Carrega as variáveis de ambiente
	config.LoadEnv()

	// Lê a porta do servidor a partir do .env
	port := config.GetEnv("APP_PORT", "8080")

	// Inicializa o banco de dados
	services.InitDB()

	// Verifica se a conexão foi bem-sucedida
	if services.DB == nil {
		log.Fatal("Erro: Banco de dados não inicializado")
	}

	// Inicializa o roteador
	router := mux.NewRouter()

	// Define uma rota de exemplo
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bem-vindo ao Portal da ONG!"))
	})

	// Registra as rotas de usuários
	routes.RegisterUsersRoutes(router)

	// Inicia o servidor
	log.Printf("Servidor iniciado na porta %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
