package main

import "fmt"

type product struct {
	id    int
	title string
	price float64
}

func main() {
	// dynamicListOfProducts slice
	dynamicListOfProducts := []product{
		{
			id:    1,
			title: "Laptop",
			price: 1000,
		},
		{
			id:    2,
			title: "Smartphone",
			price: 500,
		},
	}

	fmt.Printf("Dynamic Array of Products : %v", dynamicListOfProducts)

	dynamicListOfProducts = append(dynamicListOfProducts, product{
		id:    3,
		title: "Tablet",
		price: 700,
	})

	fmt.Printf("\nUpdated Dynamic Array of Products : %v", dynamicListOfProducts)
}

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
