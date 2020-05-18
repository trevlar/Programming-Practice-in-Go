package main

import (
	"fmt"
	"math"
)

func main() {
	books := []int{0, 1, 2, 3, 0, 1, 2, 0, 1, 0, 1, 1}
	price := calculatePriceOfBooks(books)
	expected := (4 * 8 * 0.80) +
		(3 * 8 * 0.90) +
		(4 * 8 * 0.95) +
		(1 * 8)
	if formatPrice(price) != formatPrice(expected) {
		fmt.Errorf("Price from %v was incorrect, got: %s, want: %s.", books, formatPrice(price), formatPrice(expected))
	}
}

func calculatePriceOfBooks(books []int) float64 {
	priceInFive := getDistinctPricesBasedOnStartSize(books, 4)
	priceInFour := getDistinctPricesBasedOnStartSize(books, 5)

	if priceInFive > 0 && priceInFive < priceInFour {
		return priceInFive
	}

	return priceInFour

}

func getDistinctPricesBasedOnStartSize(books []int, size int) float64 {
	totalPrice := 0.0
	booksRemainToPrice := len(books) > 0

	for booksRemainToPrice {

		// If I can acquire groups of the size passed in use it otherwise default to distinct division.
		booksInGroup, groupOfCount := getGroupsOfDiscountedBooks(books, size)
		if groupOfCount > 0 {
			discountRate, err := getDiscountRate(size)
			if err != nil {
				break
			}

			books = removeBooksFromList(books, booksInGroup)
			booksRemainToPrice = len(books) > 0

			totalPrice += float64(groupOfCount) * (8 * float64(size) * discountRate)
		} else {

			distinctBooks := getDistinctBooks(books, nil)

			discountRate, err := getDiscountRate(len(distinctBooks))
			if err != nil {
				break
			}

			if discountRate == 0 {
				totalPrice += float64(len(books)) * 8
				booksRemainToPrice = false
			} else {
				totalPrice += float64(len(distinctBooks)) * 8 * discountRate
				books = removeBooksFromList(books, distinctBooks)
				booksRemainToPrice = len(books) > 0
			}
		}
	}
	return totalPrice
}

func getDiscountRate(bookCount int) (discountRate float64, err error) {
	switch bookCount {
	case 0:
		err = fmt.Errorf("no more books to price")
	case 1:
		discountRate = 0
	case 2:
		discountRate = 0.95
	case 3:
		discountRate = 0.90
	case 4:
		discountRate = 0.80
	case 5:
		discountRate = 0.75
	default:
		err = fmt.Errorf("invalid book number")
	}
	return discountRate, err
}

func getGroupsOfDiscountedBooks(books []int, size int) (results []int, groupCount int) {

	var booksRemainToGroup = true
	booksRemaining := books
	for booksRemainToGroup {

		distinctBooks := getDistinctBooks(booksRemaining, &size)
		booksRemaining = removeBooksFromList(booksRemaining, distinctBooks)

		previousResultCount := len(results)
		results = append(results, distinctBooks...)

		booksRemainToGroup = len(booksRemaining) >= size && len(results) > previousResultCount
	}
	groupCount = len(results) / size

	return results, groupCount
}

func getDistinctBooks(books []int, size *int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range books {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
			if size != nil && len(list) == *size {
				break
			}
		}
	}
	return list
}

func removeBooksFromList(books []int, toRemove []int) []int {
	var result []int
	result = append(result, books...)
	for _, removedBook := range toRemove {
		for b, book := range result {
			if book == removedBook {
				result = remove(result, b)
				break
			}
		}
	}
	return result
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func formatPrice(price float64) string {
	return fmt.Sprintf("$%4.2f", math.Round(price*100)/100)
}
