package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	// Construir string de conexão com base nas variáveis de ambiente
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar ao banco de dados
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Verificar a conexão
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao verificar conexão com o banco de dados: %v", err)
	}

	fmt.Println("HeartBeat: MYSQL connected succefully")

	return db
}
