package types

import (
	"context"

	"github.com/QianqianZhou1214/kitchen-microservice-grpc/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
