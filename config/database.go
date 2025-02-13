package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Configuração de conexão com o PostgreSQL
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
)

// ConnectDatabase inicializa a conexão com o banco de dados
func ConnectDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
		return nil, err
	}

	// Testa a conexão
	if err := db.Ping(); err != nil {
		log.Fatal("Banco de dados não acessível:", err)
		return nil, err
	}

	fmt.Println("Conexão com PostgreSQL estabelecida!")
	return db, nil
}
