package main

import "testing"

// Test 1: AddItem
func TestAddItem(t *testing.T) {

	items = []Item{}

	err := AddItem(1, "Laptop", 5, 50000)

	if err != nil {
		t.Errorf("AddItem failed: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}
}

// Test 2: UpdateItem
func TestUpdateItem(t *testing.T) {

	items = []Item{
		{ID: 1, Name: "Laptop", Quantity: 5, Price: 50000},
	}

	err := UpdateItem(1, 10)

	if err != nil {
		t.Errorf("UpdateItem failed: %v", err)
	}

	if items[0].Quantity != 10 {
		t.Errorf("Expected quantity 10, got %d", items[0].Quantity)
	}
}

// Test 3: RemoveItem
func TestRemoveItem(t *testing.T) {

	items = []Item{
		{ID: 1, Name: "Laptop", Quantity: 5, Price: 50000},
	}

	err := RemoveItem(1)

	if err != nil {
		t.Errorf("RemoveItem failed: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

// Test 4: Input validation
func TestInvalidItem(t *testing.T) {

	items = []Item{}

	err := AddItem(0, "Laptop", 5, 50000)

	if err == nil {
		t.Error("Expected error for invalid ID")
	}
}
