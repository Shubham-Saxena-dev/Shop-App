package services

import (
	"Shop-app/proto/stubs/order"
	"context"
)

type OrderService struct{}

func (order *OrderService) GetOrdersForUser(ctx context.Context, request *order.OrderRequest) (*order.OrderResponse, error){
	panic("implement me")
}
