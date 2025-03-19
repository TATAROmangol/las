package grpcv1

import (
	"context"
	"fmt"
	"lms/internal/entities"
	ssov1 "lms/pkg/api/test/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrorInvalidId       = status.Error(codes.InvalidArgument, "invalid id")
	ErrorInvalidItem     = status.Error(codes.InvalidArgument, "invalid item")
	ErrorInvalidQuantity = status.Error(codes.InvalidArgument, "invalid quantity")
)

func getStorageError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("failed in storage: %v", err))
}

type serverAPI struct {
	ssov1.UnimplementedOrderServiceServer
	service Servicer
}

func RegisterService(gRPCServer *grpc.Server, service Servicer) {
	ssov1.RegisterOrderServiceServer(gRPCServer, &serverAPI{service: service})
}

func ValidateOrder(order entities.Order) ssov1.Order {
	return ssov1.Order{Id: order.Id, Item: order.Item, Quantity: int32(order.Quantity)}
}

func (s *serverAPI) CreateOrder(ctx context.Context, req *ssov1.CreateOrderRequest) (*ssov1.CreateOrderResponse, error) {
	if req.Item == "" {
		return nil, ErrorInvalidItem
	}

	if req.Quantity == 0 {
		return nil, ErrorInvalidQuantity
	}

	id := s.service.CreateOrder(req.GetItem(), int(req.GetQuantity()))
	return &ssov1.CreateOrderResponse{Id: id}, nil
}

func (s *serverAPI) DeleteOrder(ctx context.Context, req *ssov1.DeleteOrderRequest) (*ssov1.DeleteOrderResponse, error) {
	if len(req.Id) < 36 {
		return nil, ErrorInvalidId
	}

	resp, err := s.service.DeleteOrder(req.GetId())
	if err != nil {
		return &ssov1.DeleteOrderResponse{Success: resp}, getStorageError(err)
	}

	return &ssov1.DeleteOrderResponse{Success: resp}, nil
}

func (s *serverAPI) GetOrder(ctx context.Context, req *ssov1.GetOrderRequest) (*ssov1.GetOrderResponse, error) {
	if len(req.Id) < 36 {
		return nil, ErrorInvalidId
	}

	order, err := s.service.GetOrder(req.GetId())
	if err != nil {
		return nil, getStorageError(err)
	}

	res := ValidateOrder(order)

	return &ssov1.GetOrderResponse{Order: &res}, nil
}

func (s *serverAPI) ListOrders(ctx context.Context, req *ssov1.ListOrdersRequest) (*ssov1.ListOrdersResponse, error) {
	orders := s.service.ListOrders()

	res := make([]*ssov1.Order, len(orders))
	for i, temp := range orders {
		order := ValidateOrder(temp)
		res[i] = &order
	}

	return &ssov1.ListOrdersResponse{Orders: res}, nil
}

func (s *serverAPI) UpdateOrder(ctx context.Context, req *ssov1.UpdateOrderRequest) (*ssov1.UpdateOrderResponse, error) {
	if len(req.Id) < 36 {
		return nil, ErrorInvalidId
	}

	if req.Item == "" {
		return nil, ErrorInvalidItem
	}

	if req.Quantity == 0 {
		return nil, ErrorInvalidQuantity
	}

	order, err := s.service.UpdateOrder(req.GetId(), req.GetItem(), int(req.GetQuantity()))
	if err != nil {
		return nil, getStorageError(err)
	}

	res := ValidateOrder(order)

	return &ssov1.UpdateOrderResponse{Order: &res}, nil
}
