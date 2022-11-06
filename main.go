package main

import (
	"fmt"
)

// Car struct
type Car struct {
	color string
	brand string
}

// Product struct
type Product struct {
	id         int
	product    Car
	inStockQty int
	price      float64
}

// Display product
func (product Product) display() {
	fmt.Println("Product ID: ", product.id)
	fmt.Println("Product: ", product.product)
	fmt.Println("Quantity Available: ", product.inStockQty)
	fmt.Println("Price: ", product.price)
}

// Is product in stock or not
func (product Product) isInStock() bool {
	if product.inStockQty == 0 {
		return false
	}

	return true
}

// Store struct
type Store struct {
	products  []*Product
	soldItems []Product
}

// List all products
func (store Store) listProducts() {
	fmt.Println()
	fmt.Println("List of products:\n")
	for _, product := range store.products {
		product.display()
		fmt.Println()
	}
}

// Quantity of products available for sale
func (store Store) qtyOfProductsForSale() {
	qty := 0
	for _, product := range store.products {
		if product.isInStock() {
			qty += 1
		}
	}

	fmt.Println(qty, "products are available for sale\n")
}

// Add product to store
func (store *Store) addProduct(product *Product) {
	fmt.Println("Adding product:", product)
	store.products = append(store.products, product)
}

// Sell product in store
func (store *Store) sellProduct(id int) {
	fmt.Println("Sell product which has the id of:", id)
	for _, product := range store.products {
		if product.id == id {
			product.inStockQty -= 1
			store.soldItems = append(store.soldItems, *product)
		}
	}
}

// Display sold items
func (store Store) displaySoldItems() {
	fmt.Println("List of sold items:")
	totalPrice := 0.0
	for _, item := range store.soldItems {
		item.display()
		totalPrice += item.price
	}
	fmt.Println("Total of sold items: ", totalPrice)
}

func main() {
	productOne := Product{
		id:         1,
		product:    Car{color: "blue", brand: "toyota"},
		inStockQty: 5,
		price:      20000,
	}

	productTwo := Product{
		id:         2,
		product:    Car{color: "red", brand: "lexus"},
		inStockQty: 5,
		price:      90000,
	}

	productThree := Product{
		id:         3,
		product:    Car{color: "green", brand: "bmw"},
		inStockQty: 5,
		price:      50000,
	}

	myStore := Store{}

	myStore.addProduct(&productOne)
	myStore.addProduct(&productTwo)
	myStore.addProduct(&productThree)

	myStore.listProducts()

	myStore.sellProduct(2)
	myStore.sellProduct(2)

	myStore.listProducts()

	myStore.displaySoldItems()

	myStore.qtyOfProductsForSale()
}
