package models

type Order struct {
	Id        string
	UserId    string
	Item      string
	Adress    string
	Count     int
	CreatedAt int
	UpdatedAt int
}
