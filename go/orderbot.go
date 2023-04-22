package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var itemPool = [...]struct {
	id       string
	price    int
	category string
}{
	{"burger", 2000, "food"},
	{"schnitzel", 2100, "food"},
	{"capuns", 2800, "food"},
	{"salad", 1100, "food"},
	{"fries", 800, "food"},
	{"coke", 600, "drinks"},
	{"lemonade", 600, "drinks"},
	{"shorley", 550, "drinks"},
	{"water", 450, "drinks"},
	{"ice-cream", 700, "food"},
	{"cake", 640, "food"},
	{"pizza", 2500, "food"},
	{"pasta", 1800, "food"},
	{"steak", 3500, "food"},
	{"soup", 1200, "food"},
	{"bread", 500, "food"},
	{"coffee", 800, "drinks"},
	{"tea", 600, "drinks"},
	{"juice", 700, "drinks"},
	{"beer", 1000, "drinks"},
	{"wine", 2000, "drinks"},
	{"cheese", 1200, "food"},
}

type OrderItem struct {
	ID       string `json:"id"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}

type Order struct {
	TableNumber int         `json:"tableNumber"`
	Items       []OrderItem `json:"items"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Start sending orders
	for {
		order := generateOrder()
		sendOrder(order)
		printOrderSummary(order)
		sleepNormal()
	}
}

func generateOrder() Order {
	tableNumber := 1 + rand.Intn(9)
	numItems := 1 + poisson(1)

	// Choose random order items from the item pool
	items := make([]OrderItem, 0)
	for i := 0; i < numItems; i++ {
		items = append(items, generateOrderItem())
	}

	return Order{
		TableNumber: tableNumber,
		Items:       items,
	}
}

func generateOrderItem() OrderItem {
	item := itemPool[rand.Intn(len(itemPool))]

	// Generate a random quantity between 1 and 5
	quantity := rand.Intn(5) + 1

	return OrderItem{
		ID:       item.id,
		Price:    item.price,
		Category: item.category,
		Quantity: quantity,
	}
}

// poisson generates a random number drawn from a Poisson distribution with
// parameter lambda using the Knuth algorithm.
func poisson(lambda float64) int {
	L := math.Exp(-lambda)
	k := 0.0
	p := 1.0

	for p >= L {
		k++
		p *= rand.Float64()
	}

	return int(k - 1)
}

func sendOrder(order Order) error {
	jsonBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:3000/orders", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return nil
}

func printOrderSummary(order Order) {
	var totalCost float64
	for _, item := range order.Items {
		itemCost := float64(item.Price) * float64(item.Quantity) / 100.0
		totalCost += itemCost
	}

	fmt.Printf("Order sent: Table %d - Total: CHF %.2f\n", order.TableNumber, totalCost)
}

func sleepNormal() {
	mean := 1.0   // mean in seconds
	stdDev := 0.5 // standard deviation in seconds

	// create a normal distribution with the given mean and standard deviation
	dist := rand.New(rand.NewSource(time.Now().UnixNano())).NormFloat64

	// generate a random value from the distribution and scale it to seconds
	sleepTime := time.Duration(dist()*stdDev+mean) * time.Second

	time.Sleep(sleepTime)
}
