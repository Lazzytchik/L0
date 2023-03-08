package models

import (
	"errors"
	"fmt"
)

type Order struct {
	OrderUid          string `json:"order_uid"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	Delivery          `json:"delivery"`
	Payment           `json:"payment"`
	Items             []Item `json:"items"`
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerID        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	ShardKey          string `json:"shardkey"`
	SMID              int    `json:"sm_id"`
	DateCreated       int64  `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}

func (d Order) TableName() string {
	return "orders"
}

func (d Order) PrimaryColumn() string {
	return "id"
}

func (d Order) Validate(err []error) bool {
	switch {
	case d.OrderUid == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid OrderUid given: %s", d.OrderUid)))
		return false
	case d.TrackNumber == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid TrackNumber given: %s", d.TrackNumber)))
		return false
	case d.Entry == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Entry given: %s", d.Entry)))
		return false
	case !d.Delivery.Validate(err):
		return false
	case !d.Payment.Validate(err):
		return false
	case d.Locale == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Locale given: %s", d.Locale)))
		return false
	case d.InternalSignature == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid Internal signature given: %s", d.InternalSignature)))
		return false
	case d.CustomerID == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid CustomerID given: %s", d.CustomerID)))
		return false
	case d.DeliveryService == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid DeliverService given: %s", d.DeliveryService)))
		return false
	case d.ShardKey == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid ShardKey given: %s", d.ShardKey)))
		return false
	case d.SMID == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid SMID given: %d", d.SMID)))
		return false
	case d.DateCreated == 0:
		err = append(err, errors.New(fmt.Sprintf("Invalid DateCreated given: %d", d.DateCreated)))
		return false
	case d.OofShard == "":
		err = append(err, errors.New(fmt.Sprintf("Invalid OofShard given: %s", d.OofShard)))
		return false
	}

	for _, v := range d.Items {
		if !v.Validate(err) {
			return false
		}
	}

	return true
}
