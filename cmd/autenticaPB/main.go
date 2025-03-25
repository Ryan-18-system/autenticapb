package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ryan-18-system/autenticaPB/configs"
	"github.com/Ryan-18-system/autenticaPB/internal/infra/database"
	"github.com/Ryan-18-system/autenticaPB/internal/infra/webserver/handlers"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	// Inicializa o banco de dados
	db := database.GetDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Erro ao fechar o banco de dados:", err)
		}
	}()
	server := http.NewServeMux()
	server.HandleFunc("/user", handlers.CreateUser)
	log.Println("Servidor rodando na porta 8080...")
	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.WebServerPort), server)
	if err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}

}
