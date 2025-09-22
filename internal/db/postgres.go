package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Config pode ser extendido se quiser parametrizar host, user, pass
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB cria e retorna uma conexão com o Postgres
func NewPostgresDB() (*sql.DB, error) {
	cfg := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1234",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	// Testa a conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return db, nil
}
