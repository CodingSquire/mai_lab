package api

import (
	context "context"
	"orders/internal/controllers"
	"orders/internal/dtos"
	"orders/internal/models"
	"orders/rpc/orders"

	"github.com/lucsky/cuid"
)

type TwirpServer struct{
	OrderController controllers.OrderController
}

// CreateOrder creates a new order
func (s *TwirpServer) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := models.Order {
		ID: cuid.New(),
		UserID: req.UserId,
		Item: req.Item,
		Count: int(req.Count),
		Address: req.Adress,
	}

	err := s.OrderController.PostOrder(&order)
	twirpOrder := dtos.TwirpFromOrder(&order)

	return &orders.CreateOrderResponse{ Order: twirpOrder }, err
}

// DeleteOrder deletes an order
func (s *TwirpServer) DeleteOrder(ctx context.Context, req *orders.DeleteOrderRequest) (*orders.DeleteOrderResponse, error) {
	err := s.OrderController.DeleteOrderById(req.Id)
	return &orders.DeleteOrderResponse{}, err
}

// GetAllOrders gets all orders
func (s *TwirpServer) GetAllOrders(ctx context.Context, req *orders.GetAllOrdersRequest) (*orders.GetAllOrdersResponse, error) {
	got_orders := s.OrderController.GetAllOrders()
	twirpOrders := dtos.TwirpFromOrders(got_orders)

	return &orders.GetAllOrdersResponse{ Orders: twirpOrders }, nil
}

// GetAllOrdersByUserId gets all orders by user id
func (s *TwirpServer) GetAllOrdersByUserId(ctx context.Context, req *orders.GetAllOrdersByUserIdRequest) (*orders.GetAllOrdersByUserIdResponse, error) {
	got_orders := s.OrderController.GetAllOrdersByUserId(req.UserId)
	twirpOrders := dtos.TwirpFromOrders(got_orders)

	return &orders.GetAllOrdersByUserIdResponse{ Orders: twirpOrders }, nil
}

// GetOrder gets an order by id
func (s *TwirpServer) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	order, err := s.OrderController.GetOrderById(req.Id)
	twirpOrder := dtos.TwirpFromOrder(order)

	return &orders.GetOrderResponse{ Order: twirpOrder }, err
}

// UpdateOrder updates an order by id
func (s *TwirpServer) UpdateOrder(ctx context.Context, req *orders.UpdateOrderRequest) (*orders.UpdateOrderResponse, error) {
	order := models.Order {
		UserID: req.Order.UserId,
		Item: req.Order.Item,
		Count: int(req.Order.Count),
		Address: req.Order.Adress,
	}

	err := s.OrderController.PatchOrderById(req.Order.Id, &order)
	twirpOrder := dtos.TwirpFromOrder(&order)

	return &orders.UpdateOrderResponse{ Order: twirpOrder }, err
}

// NewTwirpServer creates a new TwirpServer
func NewTwirpServer(controller controllers.OrderController) orders.Orders {
	return &TwirpServer{
		OrderController: controller,
	}
}
