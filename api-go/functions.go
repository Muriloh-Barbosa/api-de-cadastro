package main

import (
	"api-go/structs"
	"api-go/utils"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func trataCadastro(usr string, m *structs.Informacoes) error {
	// lógica para tratar o cadastro
	log.Printf("Iniciando validação:")
	// salvar os dados em um banco de dados ou arquivo
	db, err := initDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	log.Println("Abrindo o banco de dados para tratamento do cadastro")

	var count int

	// verifica se o CPF já existe no banco de dados
	err = db.QueryRow("SELECT COUNT(*) FROM cadastro WHERE cpf = ?", m.Cpf).Scan(&count)
	if err != nil {
		log.Println("Erro ao consultar CPF:", err)
		return err
	}

	log.Println("Validando se o CPF já existe no banco de dados:", count)

	if count > 0 {
		// Retorna erro, pois não deve atualizar
		log.Println("CPF já cadastrado.")
		return fmt.Errorf("CPF já cadastrado")
	}

	_, err = db.Exec(`
		INSERT INTO cadastro (
			nomeCompleto, cpf, nascimento, endereco, cidade, estado, 
			telefone, email, codigo
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		m.Nomecompleto, m.Cpf, m.Nascimento, m.Endereco, m.Cidade,
		m.Estado, m.Telefone, m.Email, m.Codigo)
	if err != nil {
		log.Println("Erro ao inserir cadastro:", err)
		return err
	}

	log.Println("Cadastro inserido com sucesso!")
	utils.LogMonitor(utils.Debug, usr, "Cadastro tratado: "+m.Nomecompleto)
	return nil

}

func atualizaCadastro(usr string, m *structs.Informacoes) error {
	log.Printf("Iniciando atualização de cadastro")
	db, err := initDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	log.Println("Abrindo o banco de dados para atualização")

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM cadastro WHERE cpf = ?", m.Cpf).Scan(&count)
	if err != nil {
		log.Println("Erro ao consultar CPF:", err)
		return err
	}

	if count == 0 {
		log.Println("CPF não encontrado. Requisição PUT não pode criar.")
		return fmt.Errorf("CPF não encontrado")
	}

	_, err = db.Exec(`
		UPDATE cadastro SET 
			nomeCompleto=?, nascimento=?, endereco=?, cidade=?, estado=?, 
			telefone=?, email=?, codigo=?
		WHERE cpf=?`,
		m.Nomecompleto, m.Nascimento, m.Endereco, m.Cidade, m.Estado,
		m.Telefone, m.Email, m.Codigo, m.Cpf)
	if err != nil {
		log.Println("Erro ao atualizar cadastro:", err)
		return err
	}

	log.Println("Cadastro atualizado com sucesso!")
	utils.LogMonitor(utils.Debug, usr, "Cadastro atualizado: "+m.Nomecompleto)
	return nil
}
