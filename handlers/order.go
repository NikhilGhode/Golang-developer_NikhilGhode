package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tastybites/data"
	"tastybites/models"
)

type OrderRequest struct {
    UserID      int   `json:"user_id"`
    TableID     int   `json:"table_id"`
    MenuItemIDs []int `json:"menu_item_ids"`
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
        return
    }

    var req OrderRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON body", http.StatusBadRequest)
        return
    }

    table, ok := data.Tables[req.TableID]
    if !ok || !table.IsReserved || table.ReservedBy == nil || table.ReservedBy.ID != req.UserID {
        http.Error(w, "Table not reserved by this user", http.StatusForbidden)
        return
    }

    var items []models.MenuItem
    var total float64

    for _, itemID := range req.MenuItemIDs {
        item, found := data.MenuItems[itemID]
        if !found {
            http.Error(w, fmt.Sprintf("Menu item ID %d does not exist", itemID), http.StatusBadRequest)
            return
        }
        items = append(items, item)
        total += item.Price
    }

    orderID := len(data.Orders) + 1
    order := models.Order{
        ID:         orderID,
        UserID:     req.UserID,
        TableID:    req.TableID,
        Items:      items,
        TotalPrice: total,
        IsComplete: false,
    }

    data.Orders[orderID] = order

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "✅ Order #%d placed by %s at table %d. Total: ₹%.2f", order.ID, table.ReservedBy.Name, table.ID, total)
}
