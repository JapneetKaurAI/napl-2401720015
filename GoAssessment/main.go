package main

import (
	"errors"
	"fmt"
	"strings"
)

// Item structure
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Global slice
var items []Item

// AddItem function
func AddItem(id int, name string, quantity int, price float64) error {

	if id <= 0 {
		return errors.New("invalid item ID")
	}

	if strings.TrimSpace(name) == "" {
		return errors.New("item name cannot be empty")
	}

	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	if price < 0 {
		return errors.New("price cannot be negative")
	}

	// Check duplicate ID
	for _, item := range items {
		if item.ID == id {
			return errors.New("item already exists")
		}
	}

	newItem := Item{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}

	items = append(items, newItem)

	return nil
}

// UpdateItem function
func UpdateItem(id int, newQuantity int) error {

	if newQuantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	for i := range items {
		if items[i].ID == id {
			items[i].Quantity = newQuantity
			return nil
		}
	}

	return errors.New("item not found")
}

// RemoveItem function
func RemoveItem(id int) error {

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

// DisplayItems function
func DisplayItems() {

	fmt.Println("\n----- INVENTORY ITEMS -----")

	for _, item := range items {
		fmt.Printf(
			"ID: %d | Name: %s | Quantity: %d | Price: %.2f\n",
			item.ID,
			item.Name,
			item.Quantity,
			item.Price,
		)
	}
}

// Summary structure
type Summary struct {
	ItemCount  int
	TotalValue float64
}

// Concurrent summary using goroutine and channel
func CalculateSummary(inventory []Item) <-chan Summary {

	result := make(chan Summary)

	go func() {

		var summary Summary

		for _, item := range inventory {
			summary.ItemCount++
			summary.TotalValue += float64(item.Quantity) * item.Price
		}

		result <- summary
		close(result)

	}()

	return result
}

// Main function
func main() {

	fmt.Println("INVENTORY MANAGEMENT SYSTEM")

	// Adding items
	err := AddItem(1, "Laptop", 5, 50000)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddItem(2, "Mouse", 10, 500)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddItem(3, "Keyboard", 8, 1000)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Updating item
	err = UpdateItem(1, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item updated successfully!")
	}

	// Removing item
	err = RemoveItem(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item removed successfully!")
	}

	// Display inventory
	DisplayItems()

	// Calculate summary concurrently
	fmt.Println("\n----- INVENTORY SUMMARY -----")

	summary := <-CalculateSummary(items)

	fmt.Println("Total Items:", summary.ItemCount)
	fmt.Printf("Total Inventory Value: Rs. %.2f\n", summary.TotalValue)
}
