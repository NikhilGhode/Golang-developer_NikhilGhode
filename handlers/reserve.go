package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tastybites/data"
)

type ReserveRequest struct {
	UserID  int `json:"user_id"`
	TableID int `json:"table_id"`
}

func ReserveTable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ReserveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	table, exists := data.Tables[req.TableID]
	if !exists {
		http.Error(w, "Table does not exist", http.StatusNotFound)
		return
	}
	if table.IsReserved {
		http.Error(w, "Table already reserved", http.StatusConflict)
		return
	}

	table.IsReserved = true
	user := data.Users[req.UserID]
	table.ReservedBy = &user
	data.Tables[req.TableID] = table

	msg := fmt.Sprintf("Table %d successfully reserved by user %s", req.TableID, user.Name)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, msg)
}
