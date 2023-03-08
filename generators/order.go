package generators

import (
	"fmt"
	"lazzytchik/L0/models"
	"math/rand"
	"time"
)

type Order struct {
}

func (order Order) Generate(seed int) models.Order {
	itemCount := rand.Int()%3 + 1

	items := make([]models.Item, itemCount)
	for i := 0; i < itemCount; i++ {
		items[i] = models.Item{
			ChrtId:      seed,
			TrackNumber: fmt.Sprintf("%d_track", i+1),
			Price:       float32(seed),
			RID:         fmt.Sprintf("%d_rid", i+1),
			Name:        fmt.Sprintf("%d_name", i+1),
			Sale:        float32(seed),
			Size:        fmt.Sprintf("%d_size", i+1),
			TotalPrice:  float32(seed),
			NMID:        seed,
			Brand:       fmt.Sprintf("%d_brand", i+1),
			Status:      seed,
		}
	}

	return models.Order{
		OrderUid:    fmt.Sprintf("%d_uid", seed),
		TrackNumber: fmt.Sprintf("%d_track", seed),
		Entry:       fmt.Sprintf("%d_entry", seed),
		Delivery: models.Delivery{
			Name:    fmt.Sprintf("%d_name", seed),
			Phone:   fmt.Sprintf("%d_phone", seed),
			Zip:     fmt.Sprintf("%d_zip", seed),
			City:    fmt.Sprintf("%d_city", seed),
			Address: fmt.Sprintf("%d_address", seed),
			Region:  fmt.Sprintf("%d_region", seed),
			Email:   fmt.Sprintf("%d_email", seed),
		},
		Payment: models.Payment{
			Transaction:  fmt.Sprintf("%d_transaction", seed),
			RequestID:    fmt.Sprintf("%d_requestid", seed),
			Currency:     fmt.Sprintf("%d_currency", seed),
			Provider:     fmt.Sprintf("%d_provider", seed),
			Amount:       float32(seed),
			PaymentDT:    rand.Int63(),
			Bank:         fmt.Sprintf("%d_bank", seed),
			DeliveryCost: float32(seed),
			GoodsTotal:   seed,
			CustomFee:    float32(seed),
		},
		Items:             items,
		Locale:            fmt.Sprintf("%d_locale", seed),
		InternalSignature: fmt.Sprintf("%d_internalsig", seed),
		CustomerID:        fmt.Sprintf("%d_customerid", seed),
		DeliveryService:   fmt.Sprintf("%d_deliveryservice", seed),
		ShardKey:          fmt.Sprintf("%d_shardkey", seed),
		SMID:              seed,
		DateCreated:       time.Now().Unix(),
		OofShard:          fmt.Sprintf("%d_oofshard", seed),
	}
}
