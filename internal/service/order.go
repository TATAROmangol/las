package service

import (
	"lms/internal/entities"
)

type Repo interface{
	CreateOrder(string, int) string
	GetOrder(string) (entities.Order, error)
	UpdateOrder(string, string, int) (entities.Order, error)
	DeleteOrder(string) (bool, error)
	ListOrders() []entities.Order
}

type OrderService struct{
	repo Repo
}

func NewOrderService(repo Repo) *OrderService{
	return &OrderService{repo}
}

func (os *OrderService) CreateOrder(item string, quantity int) string {
	return os.repo.CreateOrder(item, quantity)
}

func (os *OrderService) DeleteOrder(id string) (bool, error) {
	return os.repo.DeleteOrder(id)
}

func (os *OrderService) GetOrder(id string) (entities.Order, error) {
	return os.repo.GetOrder(id)
}

func (os *OrderService) ListOrders() []entities.Order {
	return os.repo.ListOrders()
}

func (os *OrderService) UpdateOrder(id, item string, quantity int) (entities.Order, error) {
	return os.repo.UpdateOrder(id, item, quantity)
}

