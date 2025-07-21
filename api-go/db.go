package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func initDB() (*sql.DB, error) {
	// Carrega variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o .env: %v", err)
	}

	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	porta := os.Getenv("DB_PORT")
	nomeBanco := os.Getenv("DB_NAME")

	// Monta a string de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", usuario, senha, host, porta, nomeBanco)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Conexão com MySQL estabelecida!")

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS cadastro (
        cpf BIGINT PRIMARY KEY,
        nomeCompleto VARCHAR(100),
        nascimento DATE,
        endereco VARCHAR(150),
        cidade VARCHAR(100),
        estado VARCHAR(20),
        telefone VARCHAR(15),
        email VARCHAR(100),
        codigo VARCHAR(50)
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
