package entities

type Client struct {
	ID      int `json:"id"`
	Balance int `json:"saldo"`
	Limit   int `json:"limite"`
}
