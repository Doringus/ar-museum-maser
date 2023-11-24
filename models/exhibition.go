package models

type Exhibition struct {
	ID uint64 `gorm:"primaryKey"`
	Name string `json:"name"`
	ClientId uint64 `json:"clientId"`
	Client Client `gorm:"foreignKey:ClientId"`
}
