package domain

import "time"

// Profile representa a estrutura da tabela profiles.
type Profile struct {
	ID                    int       `json:"id"`
	UserID                int       `json:"user_id"`
	FullName              string    `json:"full_name,omitempty"`
	Address               string    `json:"address,omitempty"`
	PhoneNumber           string    `json:"phone_number,omitempty"`
	MercadoLivreUserID    string    `json:"mercado_livre_user_id,omitempty"`
	MercadoLivreAccessToken string  `json:"mercado_livre_access_token,omitempty"`
	Plan                  string    `json:"plan,omitempty"`
	PaymentType           string    `json:"payment_type,omitempty"`
	DueDate               *time.Time `json:"due_date,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
	City                  string    `json:"city,omitempty"`
	State                 string    `json:"state,omitempty"`
	BirthDate             *time.Time `json:"birthdate,omitempty"`
	Country               string    `json:"country,omitempty"`
	RefreshToken          string    `json:"refresh_token,omitempty"`
	TokenSavedAt          *time.Time `json:"token_saved_at,omitempty"`
	TokenAdjustedDate     *time.Time `json:"token_adjusted_date,omitempty"`
	AdminGroup            bool      `json:"admin_group,omitempty"`
	MetaProfile           string    `json:"meta_profile,omitempty"` 
}


// ProfilesRepository define as operações permitidas no repositório de profiles.
type ProfilesRepository interface {
	GetProfilesByUserID(userID int) ([]Profile, error)
}