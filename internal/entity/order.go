package entity

import "time"

type Order struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
