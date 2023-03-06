package storage

import "lazzytchik/L0/models"

type Orders struct {
	id      int
	storage map[int]models.Order
}

func New() Orders {
	return Orders{
		id:      1,
		storage: make(map[int]models.Order),
	}
}

func (s *Orders) Add(order models.Order) {
	s.storage[s.id] = order
	s.id++
}

func (s *Orders) GetById(id int) models.Order {
	order, _ := s.storage[id]
	return order
}

func (s *Orders) Delete(id int) {
	delete(s.storage, id)
}
