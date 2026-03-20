package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func InitDB() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := strings.Trim(os.Getenv("DB_PASSWORD"), `'"`)
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	if dbHost == "" || dbUsername == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Fatal("Missing one or more required DB_* env vars")
	}

	encodedPassword := url.QueryEscape(dbPassword)
	encodedHost := url.QueryEscape(dbHost)
	// Proper PostgreSQL connection string:
	// postgresql://user:password@host:port/dbname?sslmode=require|disable

	dbURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbUsername, encodedPassword, encodedHost, dbPort, dbName, dbSSLMode,
	)

	// cria o pool de conexões
	config, err := pgxpool.ParseConfig(dbURL)

	if err != nil {
		log.Fatalf("error with database connection %v", err)
	}

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)

	err = Pool.Ping(context.Background())

	if err != nil {
		log.Fatalf("error while connect the database %v", err)
	}

}
