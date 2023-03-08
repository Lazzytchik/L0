package storage

import (
	"errors"
	"lazzytchik/L0/models"
)

type Orders struct {
	id      int
	storage map[int]models.Order
}

func New(storage map[int]models.Order) Orders {
	return Orders{
		storage: storage,
	}
}

func (s *Orders) Add(id int, order models.Order) {
	s.storage[id] = order
}

func (s *Orders) GetById(id int) (models.Order, error) {
	if order, exists := s.storage[id]; exists {
		return order, nil
	}
	return models.Order{}, errors.New("no order in memory")
}

func (s *Orders) Delete(id int) {
	delete(s.storage, id)
}
