package models

type Exhibit struct {
	ID uint64 `gorm:"primaryKey"`
	Name string `json:"name"`
	ExhibitionId uint64 `json:"exhibitionId"`
	Exhibition Exhibition `gorm:"foreignKey:ExhibitionId"`
	ClientId uint64 `json:"clientId"`
	Client Client `gorm:"foreignKey:ClientId"`
}
