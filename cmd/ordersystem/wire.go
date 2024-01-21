//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/LeoAntunesBrombilla/clean_arch/internal/entity"
	"github.com/LeoAntunesBrombilla/clean_arch/internal/event"
	"github.com/LeoAntunesBrombilla/clean_arch/internal/infra/database"
	"github.com/LeoAntunesBrombilla/clean_arch/internal/infra/web"
	"github.com/LeoAntunesBrombilla/clean_arch/internal/usecase"
	"github.com/LeoAntunesBrombilla/clean_arch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	event.NewOrdersListed,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrderListEvent = wire.NewSet(
	event.NewOrdersListed,
	wire.Bind(new(events.EventInterface), new(*event.OrdersListed)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewGetListOfOrders(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListEvent,
		usecase.NewListOrdersUseCase,
	)

	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
