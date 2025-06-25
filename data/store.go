package data

import "tastybites/models"

var Users = map[int]models.User{
	1: {ID: 1, Name: "Nikhil"},
	2: {ID: 2, Name: "Piyush"},
}

var MenuItems = map[int]models.MenuItem{
	1: {ID: 1, Name: "Pizza", Price: 150},
	2: {ID: 2, Name: "Burger", Price: 180},
	3: {ID: 3, Name: "Pasta", Price: 120},
}

var Tables = map[int]models.Table{}
var Orders = map[int]models.Order{}

func InitTables() {
	for i := 1; i <= 20; i++ {
		Tables[i] = models.Table{
			ID:         i,
			IsReserved: false,
			ReservedBy: nil,
		}
	}
}
