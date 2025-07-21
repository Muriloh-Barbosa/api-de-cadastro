package main

import (
	"api-go/structs" // Importando structs
	"api-go/utils"   // Importando utils
	"encoding/json"
	"os"
	"strings"
	"sync"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

var mutexCadastro sync.Mutex

func userHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		usr := "Chamada para o usuário"

		utils.LogMonitor(utils.Debug, usr, "iniciando")

		if !accessValidation(c.GetHeader("client-id"), c.GetHeader("client-secret")) {
			utils.LogMonitor(utils.Debug, usr, "acesso negado")
			c.JSON(403, gin.H{"message": "Acesso negado"})
			return
		}

		utils.LogMonitor(utils.Debug, usr, "validação de acesso")
		utils.LogMonitor(utils.Debug, usr, "acesso permitido")

		//  GET
		if c.Request.Method == "GET" {
			cpf := c.Query("cpf")
			if cpf == "" {
				c.JSON(400, gin.H{"error": "CPF não informado"})
				return
			}

			db, err := initDB()
			if err != nil {
				c.JSON(500, gin.H{"error": "Erro ao conectar ao banco"})
				return
			}
			defer db.Close()

			row := db.QueryRow("SELECT nomeCompleto, nascimento, endereco, cidade, estado, telefone, email, codigo FROM cadastro WHERE cpf = ?", cpf)

			var m structs.Informacoes
			m.Cpf = cpf
			err = row.Scan(&m.Nomecompleto, &m.Nascimento, &m.Endereco, &m.Cidade, &m.Estado, &m.Telefone, &m.Email, &m.Codigo)
			if err != nil {
				c.JSON(404, gin.H{"error": "Cadastro não encontrado"})
				return
			}
			c.JSON(200, gin.H{
				"message": "Cadastro encontrado com sucesso",
				"data":    m,
			})
			return
		}

		// POST E PUT
		var Cadastro structs.Informacoes

		if err := c.BindJSON(&Cadastro); err != nil {
			utils.LogMonitor(utils.ErroGeral, usr, "Erro ao fazer bind JSON: "+err.Error())
			c.JSON(400, gin.H{"error": "Erro ao processar dados"})
			return
		}

		if _, err := govalidator.ValidateStruct(Cadastro); err != nil {
			utils.LogMonitor(utils.ErroGeral, usr, "Erro de validação: "+err.Error())
			c.JSON(400, gin.H{"error": "Dados inválidos"})
			return
		}

		JSON, _ := json.MarshalIndent(Cadastro, "", "  ")
		println("Dados recebidos:", string(JSON))
		utils.LogMonitor(utils.Debug, usr, "Dados recebidos: "+string(JSON))

		mutexCadastro.Lock()
		defer mutexCadastro.Unlock()

		var err error
		if c.Request.Method == "POST" {
			err = trataCadastro(usr, &Cadastro)
		} else if c.Request.Method == "PUT" {
			err = atualizaCadastro(usr, &Cadastro)
		}

		if err != nil {
			msg := err.Error()
			utils.LogMonitor(utils.ErroGeral, usr, "Erro ao tratar cadastro: "+msg)

			if strings.Contains(msg, "CPF já cadastrado") || strings.Contains(msg, "CPF não encontrado") {
				c.JSON(409, gin.H{"error": msg})
			} else {
				c.JSON(500, gin.H{"error": "Erro interno do servidor"})
			}
			return
		}

		c.JSON(200, gin.H{"message": "Cadastro processado com sucesso"})
	}
}

func accessValidation(clientId string, clientSecret string) bool {
	if strings.Compare(clientId, os.Getenv("client_id")) == 0 &&
		strings.Compare(clientSecret, os.Getenv("client_secret")) == 0 {
		return true
	}
	return false

}
