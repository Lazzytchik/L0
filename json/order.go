package json

import (
	"encoding/json"
	"fmt"
	"lazzytchik/L0/generators"
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

	j, _ := json.Marshal(generators.Order{}.Generate(randNumber))
	return j
}
