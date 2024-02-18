# Use a imagem base do Go
FROM golang:1.21

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos necessários para o diretório de trabalho
COPY . .

# Instale o pacote do driver MySQL
RUN go get -u github.com/go-sql-driver/mysql

# Build do projeto
RUN go build -o webservice

# Exponha a porta em que o servidor Go estará escutando
EXPOSE 8001

# Comando para iniciar o aplicativo
CMD ["./webservice"]