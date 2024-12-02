package main

import (
	"github.com/sysbietec/unicampaign/infrastructure/persistence"
	"github.com/sysbietec/unicampaign/infrastructure/router"
	"github.com/sysbietec/unicampaign/internal/adapter/controllers"
	"github.com/sysbietec/unicampaign/internal/adapter/db"
	"github.com/sysbietec/unicampaign/internal/usecases"
	"github.com/sysbietec/unicampaign/internal/logger"
)

func main() {
	logger.SetupLogger()
	logger.Info("Iniciando API Uni Campanhas")

	// Configura o banco de dados
	database, err := persistence.SetupDataBase()
	if err != nil {
		logger.Error(err)
		return
	}

	// Cria o repositório
	opportunitiesRepo := &db.OpportunitiesRepositoryImpl{DB: database}

	// Cria o caso de uso
	opportunitiesUC := usecases.NewOpportunitiesUseCase(opportunitiesRepo)

	// Cria o controlador
	opportunitiesController := controllers.NewOpportunitiesController(opportunitiesUC)

	// Configura o roteador
	r := router.SetupRouter(opportunitiesController)

	// Inicia o servidor
	if err := r.Run(":8080"); err != nil {
		logger.Error(err)
		return
	}
}
