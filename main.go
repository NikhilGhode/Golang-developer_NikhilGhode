package main

import (
	"fmt"
	"net/http"
	"tastybites/data"
	"tastybites/handlers"
)

func main() {
	fmt.Println("TastyBites backend is running...")

	data.InitTables()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to TastyBites!")
	})

	http.HandleFunc("/reserve", handlers.ReserveTable)

	http.HandleFunc("/order", handlers.PlaceOrder)

	http.HandleFunc("/admin/table/", handlers.AdminViewTable)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
