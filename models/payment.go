package models

import (
	"errors"
	"fmt"
)

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

func (d Payment) Validate(err []error) bool {
	switch {
	case d.Transaction == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Transaction given: %s", d.Transaction)))
		return false
	case d.RequestID == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid RequestID given: %s", d.RequestID)))
		return false
	case d.Currency == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Currency given: %s", d.Currency)))
		return false
	case d.Provider == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Provider given: %s", d.Provider)))
		return false
	case d.Amount == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid Amount given: %f", d.Amount)))
		return false
	case d.PaymentDT == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid PaymentDT given: %d", d.PaymentDT)))
		return false
	case d.Bank == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Bank given: %s", d.Bank)))
		return false
	case d.DeliveryCost == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid DeliveryCost given: %f", d.DeliveryCost)))
		return false
	case d.GoodsTotal == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid GoodsTotal given: %d", d.GoodsTotal)))
		return false
	case d.CustomFee == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid CustomFee given: %f", d.CustomFee)))
		return false
	}

	return true
}

func (d Payment) TableName() string {
	return "payments"
}

func (d Payment) PrimaryColumn() string {
	return "id"
}
