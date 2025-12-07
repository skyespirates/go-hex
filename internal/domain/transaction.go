package domain

type Transaction struct {
	Id         int    `json:"id"`
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId"`
	Amount     int    `json:"amount"`
}
