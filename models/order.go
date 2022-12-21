package models

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

type Store struct {
	Product `json:"product" binding:"required"`
	Count   int `json:"count" binding:"required"`
}

type Stores []Store

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"uid"`
	Store  Stores `json:"products"`
}
