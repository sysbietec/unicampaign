package domain
import "time"

type Opportunity struct {
	ID           string     `json:"id"`
	Type         string     `json:"type"`
	Status       string     `json:"status"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	FinishDate   *time.Time `json:"finish_date,omitempty"`
	DeadLineDate *time.Time `json:"deadline_date,omitempty"`
	Name         string     `json:"name,omitempty"`
	BenefitsType string     `json:"benefits_type,omitempty"`
	MeliPercent  int        `json:"meli_percent,omitempty"`
	SellerPercent int       `json:"seller_percent,omitempty"`
	Finished     bool       `json:"-" gorm:"default:false"`
	ProfileID     int        `json:"profile_id"`
}

type OpportunitiesRepository interface {
	GetAvailableOpportunities() ([]Opportunity, error)
	UpdateCampaignStatus() error
	GetMercadoLivreUsers() ([]Profile, error) // Método para obter usuários do Mercado Livre
	SaveOpportunity(opportunity Opportunity) error // Método para salvar oportunidade
}