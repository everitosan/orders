package retriever

import "time"

const APIURL = "http://localhost:3333/orders/status/"

type OrderStatusResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
