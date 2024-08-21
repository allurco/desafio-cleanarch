package service

import (
	"context"

	"github.com/allurco/desafio-cleanarch/internal/infra/grpc/pb"
	"github.com/allurco/desafio-cleanarch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {

	page := int64(1)
	limit := int64(10)
	sort := "asc"

	if in != nil {
		page = in.Page
		limit = in.Limit
		sort = in.Sort
	}

	orders, err := s.ListOrdersUseCase.ListOrders(int(limit), int(page), sort)
	if err != nil {
		return nil, err
	}

	ordersSLice := make([]*pb.Order, len(orders))
	for i, order := range orders {
		ordersSLice[i] = &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
	}

	orderResponse := &pb.ListOrdersResponse{
		Orders: ordersSLice,
	}

	return orderResponse, err
}
