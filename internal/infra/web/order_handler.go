package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/allurco/desafio-cleanarch/internal/entity"
	"github.com/allurco/desafio-cleanarch/internal/usecase"
	"github.com/allurco/desafio-cleanarch/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	var page = 1
	if r.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	}

	var limit = 10
	if r.URL.Query().Get("limit") != "" {
		limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}
	sort := r.URL.Query().Get("sort")

	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository)

	orders, err := listOrders.ListOrders(limit, page, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
