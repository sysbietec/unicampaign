package db

import (
	"database/sql"
	"time"

	"github.com/sysbietec/unicampaign/internal/domain"
)

type OpportunitiesRepositoryImpl struct {
	DB *sql.DB
}

// Atualiza o status das campanhas com base na data atual.
func (r *OpportunitiesRepositoryImpl) UpdateCampaignStatus() error {
	currentTime := time.Now()

	query := `
		UPDATE opportunities
		SET finished = CASE
			WHEN finish_date < $1 THEN true
			ELSE false
		END
	`
	if _, err := r.DB.Exec(query, currentTime); err != nil {
		return err
	}
	return nil
}


// Retorna as oportunidades disponíveis.
func (r *OpportunitiesRepositoryImpl) GetAvailableOpportunities() ([]domain.Opportunity, error) {
	query := `
		SELECT id, type, status, start_date, finish_date, deadline_date, name, finished
		FROM opportunities
		WHERE finished = false
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opportunities []domain.Opportunity
	for rows.Next() {
		var opportunity domain.Opportunity
		if err := rows.Scan(
			&opportunity.ID,
			&opportunity.Type,
			&opportunity.Status,
			&opportunity.StartDate,
			&opportunity.FinishDate,
			&opportunity.DeadLineDate,
			&opportunity.Name,
			&opportunity.Finished,
		); err != nil {
			return nil, err
		}
		opportunities = append(opportunities, opportunity)
	}

	return opportunities, nil
}


func (r *OpportunitiesRepositoryImpl) GetMercadoLivreUsers() ([]domain.Profile, error) {
	query := `
		SELECT id, mercado_livre_user_id
		FROM profiles
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.Profile
	for rows.Next() {
		var user domain.Profile
		if err := rows.Scan(&user.ID, &user.MercadoLivreUserID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Salva uma oportunidade no banco de dados.
func (r *OpportunitiesRepositoryImpl) SaveOpportunity(opportunity domain.Opportunity) error {
	query := `
		INSERT INTO opportunities (id, type, status, start_date, finish_date, deadline_date, name, finished, profile_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
			type = EXCLUDED.type,
			status = EXCLUDED.status,
			start_date = EXCLUDED.start_date,
			finish_date = EXCLUDED.finish_date,
			deadline_date = EXCLUDED.deadline_date,
			name = EXCLUDED.name,
			finished = EXCLUDED.finished,
			profile_id = EXCLUDED.profile_id
	`
	_, err := r.DB.Exec(query,
		opportunity.ID,
		opportunity.Type,
		opportunity.Status,
		opportunity.StartDate,
		opportunity.FinishDate,
		opportunity.DeadLineDate,
		opportunity.Name,
		opportunity.Finished,
		opportunity.ProfileID,
	)
	return err
}
