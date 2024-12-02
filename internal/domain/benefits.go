package domain


type Benefits struct {
	ID           uint   `gorm:"primaryKey"`
	Type         string `json:"type"`
	MeliPercent  int    `json:"meli_percent"`
	SellerPercent int   `json:"seller_percent"`
}