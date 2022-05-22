//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printList(product [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(product); i++ {
		products := product[i]
		cost += products.price

		if products.name != "" {
			totalItems += 1
		}

		fmt.Println("The last item in the list is ", product[totalItems-1])
		fmt.Println("Total items are ", totalItems)
		fmt.Println("Total cost is ", cost)
	}
}

func main() {
	// Add products
	shoppingList := [4]Product{
		{"Rice", 3},
		{"Meat", 5},
		{"Oats", 10},
	}

	printList(shoppingList)
	shoppingList[3] = Product{"Bread", 2}

	printList(shoppingList)
}
