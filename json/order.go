package json

import (
	"encoding/json"
	"lazzytchik/L0/generators"
	"math/rand"
)

type Order struct {
}

func (*Order) Generate() []byte {
	randNumber := rand.Int()%10 + 1

	j, _ := json.Marshal(generators.Order{}.Generate(randNumber))
	return j
}
