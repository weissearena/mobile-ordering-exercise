package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Order struct {
	ID          int         `json:"id"`
	TableNumber int         `json:"tableNumber"`
	Timestamp   int         `json:"timestamp"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	ID       string `json:"id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Accepting orders")
	case "POST":
		handleOrder(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate total price in CHF
	totalPrice := float64(0)
	for _, item := range order.Items {
		totalPrice += float64(item.Price * item.Quantity)
	}

	// Convert total price to CHF
	totalPrice /= 100 // convert rappen to CHF

	// Print out order summary
	fmt.Printf("Order #%d @ %d\t\tTable %d\t\tCHF %.2f\n", order.ID, order.Timestamp, order.TableNumber, totalPrice)
}
