package json

import (
	"encoding/json"
	"fmt"
	"lazzytchik/L0/models"
	"math/rand"
)

type Order struct {
}

func (*Order) Generate() []byte {
	randNumber := rand.Int() % 10
	itemCount := rand.Int() % 3

	items := make([]models.Item, itemCount)
	for i := 0; i < itemCount; i++ {
		items[i] = models.Item{
			ChrtId:      randNumber,
			TrackNumber: fmt.Sprintf("%d_track", i),
			Price:       float32(randNumber),
			RID:         fmt.Sprintf("%d_rid", i),
			Name:        fmt.Sprintf("%d_name", i),
			Sale:        float32(randNumber),
			Size:        fmt.Sprintf("%d_size", i),
			TotalPrice:  float32(randNumber),
			NMID:        randNumber,
			Brand:       fmt.Sprintf("%d_brand", i),
			Status:      randNumber,
		}
	}

	j, _ := json.Marshal(models.Order{
		OrderUid:    fmt.Sprintf("%d_uid", randNumber),
		TrackNumber: fmt.Sprintf("%d_track", randNumber),
		Entry:       fmt.Sprintf("%d_entry", randNumber),
		Delivery: models.Delivery{
			Name:    fmt.Sprintf("%d_name", randNumber),
			Phone:   fmt.Sprintf("%d_phone", randNumber),
			Zip:     fmt.Sprintf("%d_zip", randNumber),
			City:    fmt.Sprintf("%d_city", randNumber),
			Address: fmt.Sprintf("%d_address", randNumber),
			Region:  fmt.Sprintf("%d_region", randNumber),
			Email:   fmt.Sprintf("%d_email", randNumber),
		},
		Payment: models.Payment{
			Transaction:  fmt.Sprintf("%d_transaction", randNumber),
			RequestID:    fmt.Sprintf("%d_requestid", randNumber),
			Currency:     fmt.Sprintf("%d_currency", randNumber),
			Provider:     fmt.Sprintf("%d_provider", randNumber),
			Amount:       float32(randNumber),
			PaymentDT:    rand.Int63(),
			Bank:         fmt.Sprintf("%d_bank", randNumber),
			DeliveryCost: float32(randNumber),
			GoodsTotal:   randNumber,
			CustomFee:    float32(randNumber),
		},
		Items:             items,
		Locale:            fmt.Sprintf("%d_locale", randNumber),
		InternalSignature: fmt.Sprintf("%d_internalsig", randNumber),
		CustomerID:        fmt.Sprintf("%d_customerid", randNumber),
		DeliveryService:   fmt.Sprintf("%d_deliveryservice", randNumber),
		ShardKey:          fmt.Sprintf("%d_shardkey", randNumber),
		SMID:              randNumber,
		DateCreated:       fmt.Sprintf("%d_datecreated", randNumber),
		OofShard:          fmt.Sprintf("%d_oofshard", randNumber),
	})
	return j
}
