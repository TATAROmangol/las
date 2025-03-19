package storage

import (
	"fmt"
	"lms/internal/entities"
	"sync"

	"github.com/google/uuid"
)

type OrderRepo struct {
	mu     *sync.RWMutex
	orders map[string]entities.Order
}

func NewOrderRepo() *OrderRepo{
	return &OrderRepo{&sync.RWMutex{}, make(map[string]entities.Order)}
}

var (
	ErrorId = fmt.Errorf("not have order with this id")
)

func (or *OrderRepo) CreateOrder(item string, quantity int) string {
	or.mu.Lock()
	defer or.mu.Unlock()

	var id string
	have := true
	for have {
		id = uuid.New().String()
		if _, ok := or.orders[id]; !ok {
			have = false
		}
	}

	order := entities.NewOrder(id, item, quantity)
	or.orders[id] = order
	return id
}

func (or *OrderRepo) GetOrder(id string) (entities.Order, error) {
	or.mu.RLock()
	defer or.mu.RUnlock()
	order, ok := or.orders[id]
	if !ok{
		return entities.Order{}, ErrorId
	}

	return order, nil
}

func (or *OrderRepo) UpdateOrder(id, item string, quantity int) (entities.Order, error) {
	or.mu.Lock()
	defer or.mu.Unlock()

	if _, ok := or.orders[id]; !ok {
		return entities.Order{}, ErrorId
	}

	order := entities.NewOrder(id, item, quantity)
	or.orders[id] = order

	return order, nil
}

func (or *OrderRepo) DeleteOrder(id string) (bool, error){
	or.mu.Lock()
	defer or.mu.Unlock()

	if _, ok := or.orders[id]; !ok{
		return false, ErrorId
	}

	delete(or.orders, id)
	return true, nil
}

func (or *OrderRepo) ListOrders() []entities.Order{
	or.mu.RLock()
	defer or.mu.RUnlock()

	res := make([]entities.Order, 0, len(or.orders))
	for _, order := range or.orders{
		res = append(res, order)
	}

	return res
}
