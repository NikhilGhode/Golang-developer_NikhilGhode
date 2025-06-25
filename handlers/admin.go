package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tastybites/data"
	"tastybites/models"
)

type AdminResponse struct {
    TableID     int               `json:"table_id"`
    User        *models.User      `json:"user,omitempty"`
    Order       *models.Order     `json:"order,omitempty"`
    TotalAmount float64           `json:"total_amount"`
    Message     string            `json:"message,omitempty"`
}

func AdminViewTable(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extract table ID from URL path like /admin/table/4
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) != 4 {
        http.Error(w, "Invalid path. Use /admin/table/{id}", http.StatusBadRequest)
        return
    }

    tableID, err := strconv.Atoi(parts[3])
    if err != nil {
        http.Error(w, "Invalid table ID", http.StatusBadRequest)
        return
    }

    table, exists := data.Tables[tableID]
    if !exists {
        http.Error(w, "Table does not exist", http.StatusNotFound)
        return
    }

    if !table.IsReserved || table.ReservedBy == nil {
        resp := AdminResponse{
            TableID: tableID,
            Message: "Table is currently not reserved.",
        }
        json.NewEncoder(w).Encode(resp)
        return
    }

    // Find the latest order for this table
    var latest *models.Order
    for _, order := range data.Orders {
        if order.TableID == tableID && order.UserID == table.ReservedBy.ID {
            latest = &order
        }
    }

    resp := AdminResponse{
        TableID:     tableID,
        User:        table.ReservedBy,
        Order:       latest,
        TotalAmount: func() float64 {
            if latest != nil {
                return latest.TotalPrice
            }
            return 0
        }(),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

