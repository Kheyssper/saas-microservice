package main

import (
	"log"
	"os"
	"saasmicroservice/pkg/db"
	"saasmicroservice/pkg/router"
	"saasmicroservice/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Conectar ao banco de dados
	dbConn, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Configurar o roteador
	r := router.SetupRouter(dbConn)

	// Iniciar o servidor
	srv := server.Server{}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}
	if err := srv.Run(port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
