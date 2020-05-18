package main

import (
	"testing"
)

func TestBasics(t *testing.T) {
	var books []int
	var price float64
	var expected float64

	books = []int{}
	price = calculatePriceOfBooks(books)
	expected = 0.0
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	expected = 8.0

	books = []int{1}
	price = calculatePriceOfBooks(books)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{2}
	price = calculatePriceOfBooks(books)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{3}
	price = calculatePriceOfBooks(books)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{4}
	price = calculatePriceOfBooks(books)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}
}

func TestSimpleDiscounts(t *testing.T) {
	var books []int
	var price float64
	var expected float64

	books = []int{0, 1}
	price = calculatePriceOfBooks(books)
	expected = 8 * 2 * 0.95
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{0, 2, 4}
	price = calculatePriceOfBooks(books)
	expected = 8 * 3 * 0.9
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{0, 1, 2, 4}
	price = calculatePriceOfBooks(books)
	expected = 8 * 4 * 0.8
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{0, 1, 2, 3, 4}
	price = calculatePriceOfBooks(books)
	expected = 8 * 5 * 0.75
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}
}

func TestSeveralDiscounts(t *testing.T) {
	var books []int
	var price float64
	var expected float64

	books = []int{0, 0, 1}
	price = calculatePriceOfBooks(books)
	expected = 8 + (8 * 2 * 0.95)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{0, 0, 1, 1}
	price = calculatePriceOfBooks(books)
	expected = 2 * (8 * 2 * 0.95)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{0, 0, 1, 2, 2, 3}
	price = calculatePriceOfBooks(books)
	expected = (8 * 4 * 0.8) + (8 * 2 * 0.95)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{0, 1, 1, 2, 3, 4}
	price = calculatePriceOfBooks(books)
	expected = 8 + (8 * 5 * 0.75)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}
}

func TestEdgeCases(t *testing.T) {
	var books []int
	var price float64
	var expected float64

	books = []int{0, 0, 1, 1, 2, 2, 3, 4}
	price = calculatePriceOfBooks(books)
	expected = 2 * (8 * 4 * 0.8)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}

	books = []int{
		0, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3, 3,
		4, 4, 4, 4,
	}
	price = calculatePriceOfBooks(books)
	expected = 4*(8*5*0.75) + (8 * 3 * 0.9)
	if formatPrice(price) != formatPrice(expected) {
		t.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}
}
