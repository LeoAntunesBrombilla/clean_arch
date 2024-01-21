package usecase

import (
	"github.com/LeoAntunesBrombilla/clean_arch/internal/entity"
	"github.com/LeoAntunesBrombilla/clean_arch/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface, OrdersListed events.EventInterface, EventDispatcher events.EventDispatcherInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.GetListOfOrders()
	if err != nil {
		return nil, err
	}

	var dtos []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dtos = append(dtos, dto)
	}

	payload := map[string]interface{}{
		"totalOrders": len(dtos),
	}

	l.OrdersListed.SetPayload(payload)
	l.EventDispatcher.Dispatch(l.OrdersListed)

	return dtos, nil
}
