package controllers

import (
	"orders/internal/db"
	"orders/internal/dotenv"
	"orders/internal/models"
	"testing"

	"github.com/lucsky/cuid"
)

func TestPGXOrderController(t *testing.T) {
	dotenv.Config()
	db, err := db.Init()
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	controller := NewPgxController(db)

	ID := cuid.New()
	userID := cuid.New()
	order := models.Order {
		ID: ID,
		UserID: userID,
		Item: "CPU",
		Address: "123 Main St",
		Count: 10,
	}

	err = controller.PostOrder(&order)
	if err != nil {
		t.Errorf("Error posting order: %v", err)
	}

	got_order, err := controller.GetOrderById(ID)
	if err != nil {
		t.Errorf("Error getting order: %v", err)
	}
	if got_order.Item != order.Item {
		t.Errorf("Error mismatched order item, got %v, want %v", got_order.Item, order.Item)
	}

	orders := controller.GetAllOrders()
	if len(orders) != 1 {
		t.Errorf("Error mismatched order count, got %v, want %v, orders: %v", len(orders), 1, orders)
	}


	new_order := models.Order {
		UserID: userID,
		Item: "GPU",
		Address: "Moscov",
		Count: 1,
	}
	err = controller.PatchOrderById(ID, &new_order)
	if err != nil {
		t.Errorf("Error patching order by id")
	}

	got_order, err = controller.GetOrderById(ID)
	if err != nil {
		t.Errorf("Error getting patched order, %v", err)
	}
	if got_order.Item != new_order.Item {
		t.Errorf("Error Item mismatch, on patched order")
	}

	err = controller.DeleteOrderById(ID)
	if err != nil {
		t.Errorf("Error deleting order by id")
	}

	orders = controller.GetAllOrders()
	if len(orders) != 0 {
		t.Errorf("Error orders still exist after delete")
	}
}
