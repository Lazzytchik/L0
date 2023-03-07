package models

import "fmt"

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

func (d Payment) Insert() string {
	return fmt.Sprintf(
		"INSERT INTO %s (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ('%s', '%s', '%s', '%s', '%f', '%d', '%s', %f, %d, %f) RETURNING id",
		d.TableName(),
		d.Transaction,
		d.RequestID,
		d.Currency,
		d.Provider,
		d.Amount,
		d.PaymentDT,
		d.Bank,
		d.DeliveryCost,
		d.GoodsTotal,
		d.CustomFee,
	)
}

func (d Payment) Delete(id int) string {
	return fmt.Sprintf(
		"DELETE FROM %s WHERE %s = %d",
		d.TableName(),
		d.PrimaryColumn(),
		id,
	)
}

func (d Payment) TableName() string {
	return "payments"
}

func (d Payment) PrimaryColumn() string {
	return "id"
}
