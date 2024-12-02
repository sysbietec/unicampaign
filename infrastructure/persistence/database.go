package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/sysbietec/unicampaign/internal/logger"
	_ "github.com/lib/pq" 
)

// SetupDataBase configura a conexão com o banco de dados usando a biblioteca padrão.
func SetupDataBase() (*sql.DB, error) {
	// Cria a string de conexão com base nas variáveis de ambiente.
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable" // Define um valor padrão se a variável estiver vazia
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		sslmode,
	)

	logger.Info("Configurando conexão com o banco de dados...")
	// Abre a conexão com o banco de dados.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(fmt.Errorf("erro ao abrir conexão com o banco de dados: %w", err))
		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %w", err)
	}

	logger.Info(fmt.Sprintf("String de conexão: %s", dsn))

	// Testa a conexão.
	if err := db.Ping(); err != nil {
		logger.Error(fmt.Errorf("erro ao testar conexão com o banco de dados: %w", err))
		return nil, fmt.Errorf("erro ao testar conexão com o banco de dados: %w", err)
	}

	logger.Info("Conexão com o banco de dados configurada com sucesso.")
	return db, nil
}
