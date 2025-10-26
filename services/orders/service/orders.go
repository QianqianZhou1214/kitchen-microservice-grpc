package service

import (
	"context"

	"github.com/QianqianZhou1214/kitchen-microservice-grpc/services/common/genproto/orders"
)

// business logic here
var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	//store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}
