package grpcv1

import "lms/internal/entities"

type Servicer interface {
	CreateOrder(string, int) string
	GetOrder(string) (entities.Order, error)
	UpdateOrder(string, string, int) (entities.Order, error)
	DeleteOrder(string) (bool, error)
	ListOrders() []entities.Order
}
