package main

import (
	"aprendamais/Database"
	"aprendamais/server"
	"fmt"
)

func main() {
	// Inicia o banco de dados
	Database.StartDB()
	// Cria uma estância do server
	server := server.NewServer()
	// Inicia o server
	server.Run()

	fmt.Println("Server rodando na porta 5000")
}
