package models

type Client struct {
	ID uint64 `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

