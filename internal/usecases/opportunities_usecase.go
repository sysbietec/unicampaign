package usecases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sysbietec/unicampaign/internal/domain"
)

type OpportunitiesUseCase struct {
	Repo domain.OpportunitiesRepository
}

// NewOpportunitiesUseCase cria uma nova instância de OpportunitiesUseCase.
func NewOpportunitiesUseCase(repo domain.OpportunitiesRepository) *OpportunitiesUseCase {
	return &OpportunitiesUseCase{Repo: repo}
}

// FetchAvailableOpportunities obtém oportunidades disponíveis.
func (uc *OpportunitiesUseCase) FetchAvailableOpportunities() ([]domain.Opportunity, error) {
	// Atualiza o status das campanhas no banco de dados.
	if err := uc.Repo.UpdateCampaignStatus(); err != nil {
		return nil, err
	}

	// Busca os usuários da tabela profiles para consultar a API externa.
	users, err := uc.Repo.GetMercadoLivreUsers()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários do Mercado Livre: %w", err)
	}

	// Itera sobre os usuários e consulta a API externa para cada um.
	for _, user := range users {
		apiURL := fmt.Sprintf("https://api.mercadolibre.com/seller-promotions/users/%s?app_version=v2", user.MercadoLivreUserID)
		resp, err := http.Get(apiURL)
		if err != nil {
			return nil, fmt.Errorf("erro ao consultar a API externa: %w", err)
		}
		defer resp.Body.Close()

		// Lê e processa a resposta da API.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler resposta da API externa: %w", err)
		}

		var apiResponse struct {
			Results []domain.Opportunity `json:"results"`
		}
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			return nil, fmt.Errorf("erro ao decodificar resposta JSON da API externa: %w", err)
		}

		// Salva as oportunidades retornadas no banco de dados.
		for _, opportunity := range apiResponse.Results {
			opportunity.ProfileID = user.ID
			if err := uc.Repo.SaveOpportunity(opportunity); err != nil {
				return nil, fmt.Errorf("erro ao salvar oportunidade no banco de dados: %w", err)
			}
		}
	}

	// Retorna as oportunidades disponíveis.
	return uc.Repo.GetAvailableOpportunities()
}
