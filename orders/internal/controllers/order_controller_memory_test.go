package controllers

import (
	"github.com/innerave/mai_lab/orders/internal/models"
	"testing"

	"github.com/lucsky/cuid"
)

func TestOrderControllerMemory(t *testing.T) {
	controller := NewOrderMemController()

	id := cuid.New()
	userID := cuid.New()

	order := models.Order{
		ID:      id,
		UserID:  userID,
		Address: "",
		Item:    "",
		Count:   0,
	}

	controller.PostOrder(&order)
	got_order, err := controller.GetOrderById(id)
	if err != nil {
		t.Errorf("Error getting order by id")
	}

	_, err = controller.GetOrderById("wrong_id")
	if err == nil {
		t.Error("Error getting order")
	}

	new_order := models.Order{
		UserID:  userID,
		Address: "new_address",
		Item:    "new_item",
		Count:   1,
	}
	err = controller.PatchOrderById(id, &new_order)
	if err != nil {
		t.Error("Error updating order")
	}

	got_order, err = controller.GetOrderById(id)
	if err != nil {
		t.Error("Error getting patched order")
	}
	if got_order.Address != "new_address" {
		t.Error("Error updating order")
	}

	orders := controller.GetAllOrders()
	if len(orders) != 1 {
		t.Error("Error getting active orders")
	}
}
