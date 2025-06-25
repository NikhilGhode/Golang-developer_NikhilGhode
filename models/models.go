package models

type User struct {
	ID    int
	Name  string
	Email string
}

type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

type Table struct {
	ID         int
	IsReserved bool
	ReservedBy *User
}

type Order struct {
	ID         int
	UserID     int
	TableID    int
	Items      []MenuItem
	TotalPrice float64
	IsComplete bool
}
