package models

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestID    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       float32 `json:"amount"`
	PaymentDT    int64   `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost float32 `json:"delivery_cost"`
	GoodsTotal   int     `json:"goods_total"`
	CustomFee    float32 `json:"custom_fee"`
}
