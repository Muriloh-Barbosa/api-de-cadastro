#Validar versao do go
FROM golang:1.24.3-alpine

# Instala dependências úteis (compilação, git para go mod)
RUN apk add --no-cache git

# Cria diretório de trabalho
WORKDIR /app

# Copia os arquivos de dependência primeiro (otimiza cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código-fonte
COPY . .

# Compila o binário
RUN go build -o app

# Expõe a porta que o servidor usa
EXPOSE 8080

# Comando padrão ao rodar o container
CMD ["./app"]
