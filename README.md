# 🎯 API desenvolvida em Go com interface frontend em HTML/JavaScript, containerizada com Docker

Implementação de uma API em Go conectada a um banco de dados MySQL, acompanhada por uma interface web simples para cadastro e consulta de usuários.

A aplicação é totalmente containerizada com Docker e Docker Compose.

---

## 📌 **Funcionalidades**
- Criar novos cadastros de usuários com validação de CPF
- Atualizar cadastros existentes
- Armazenamento em banco de dados MySQL
- Interface web para interação com a API

---

## 🗂️ **Estrutura do Projeto**

api_cadastro/   
│   docker-compose.yml  
│   
├── api-go/                      # Backend em Go      
│   ├── .env                     # Configurações de ambiente    
│   ├── Dockerfile               # Imagem Docker do backend     
│   ├── main.go                  # Entrada principal da API     
│   ├── db.go                    # Conexão e inicialização do banco     
│   ├── functions.go             # Regras de negócio    
│   ├── consult.go               # Consultas GET por CPF    
│   ├── configRouter.go          # Rotas da API     
│   ├── go.mod / go.sum          # Dependências     
│   ├── structs/                 # Estruturas de dados          
│   └── utils/                   # logs    
│   
└── frontEnd/                      # Interface Web  
├── Dockerfile              # Imagem Docker do frontend 
├── index.html              # Página HTML principal 
├── css.css                 # Estilos   
└── js.js                   # interação com a API 

---

## 🛑 **Pré-Requisitos**
- Go lang
- Java script
- visual studio code
- Docker
- MySQL
- Git
- Conta no GitHub

---

## 🎯 **Visão Geral**

Backend: Desenvolvido em Go, criação (POST), (PUT) e consulta (GET). A aplicação se conecta com um banco MySQL.

Frontend: Interface HTML/CSS/JS simples e intuitiva, que interage com a API via requisições.

Banco de Dados: Container MySQL com persistência de dados.

Ambiente Dockerizado: Toda a aplicação esta em um docker-compose, facilitando a instalação e execução.

---

## 🖥️ **Execução do Projeto**

✅ **Usando o visual studio code**
1. Clone este repositório:
2. No terminal visual studio code, execute docker compose up --build
3. Necessario configurar as variaiveis de ambiente para rodar com sucesso
4. Para visualizar o teste acesse:
    - Frontend: http://localhost
    - API (teste direto): http://localhost:8080

---

## 📦 **Exemplo de interação** 
✅ Cadastrar novos usuários  
✅ Atualizar cadastros existentes  
✅ Consultar por CPF e visualizar os dados na tela  
✅ Visualizar mensagens de erro e sucesso em tempo real  

## 💡 **Possíveis Melhorias**
- Adicionar autenticação
- Implementar testes automatizados
- Criar tela para listagem completa de cadastros
- Deploy em nuvem (AWS)

---

## 👨‍💻 **Autor**
Desenvolvido por **Murilo**

✨ Transformar ideias em código. ✨

---

## 🤝 **Contribuições**
Sinta-se à vontade para abrir *issues* ou enviar *pull requests* para melhorias!
