package main

import (
	"fmt"
)

var PORT = "8080"

func main() {

	r := newRouter()

	// Iniciando
	fmt.Println("Iniciando...")

	// Iniciando o servidor
	r.Run(":" + PORT)
	fmt.Println("Servidor rodando na porta", PORT)
}
