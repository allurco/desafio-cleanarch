package usecase

import (
	"github.com/allurco/desafio-cleanarch/internal/entity"
)

type OrderListOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (l *ListOrdersUseCase) ListOrders(limit, page int, sort string) ([]OrderListOutputDTO, error) {
	orders := l.OrderRepository.List(page, limit, sort)
	orderResponse := make([]OrderListOutputDTO, len(orders))
	for i, order := range orders {
		orderResponse[i] = OrderListOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return orderResponse, nil
}
