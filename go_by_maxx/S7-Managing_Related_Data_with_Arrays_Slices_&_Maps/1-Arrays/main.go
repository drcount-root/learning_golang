package main

import "fmt"

func main() {
	// Declare an array of 4 float64 values
	prices := [4]float64{12.3, 45.6, 78.9, 23.4}

	fmt.Printf("Price array: %v, which is of length %v\n", prices, len(prices))

	// Declare an array of 6 string without initial values
	var productNames [6]string
	fmt.Printf("1 Product names array: %v, which is of length %v\n", productNames, len(productNames))
	productNames = [6]string{"Laptop", "Smartphone", "Tablet", "Monitor", "Keyboard", "Mouse"}
	fmt.Printf("2 Product names array after adding values: %v, which is of length %v\n", productNames, len(productNames))

	productNames[2] = "Smartwatch"
	fmt.Printf("3 Product names array after updating value of index 2: %v, which is of length %v\n", productNames, len(productNames))

	/*
	************* Selecting parts of array using slices *************
	 */

	names := [7]string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace"}

	namesSlice1 := names[2:5] // returns a slice without modifying the original array
	fmt.Printf("\nNames slice from index 2 to 4 excluding index 5: %v\n", namesSlice1)
	fmt.Print("\n", names)

	namesSlice2 := names[:4] // returns a slice from the start to index 3
	fmt.Printf("\nNames slice from start to index 3 excluding index 4: %v\n", namesSlice2)

	namesSlice3 := names[3:]                                               // returns a slice from index 3 to the end
	fmt.Printf("\nNames slice from index 3 to the end: %v\n", namesSlice3) // [David Eve Frank Grace]

	namesSlice3[1] = "Evelyn" // modifies the original array as well at the Eve's position
	fmt.Printf("\nnamesSlice3 after modification: %v\n", namesSlice3)
	fmt.Printf("\nOriginal names array after modifying namesSlice3: %v\n", names)

	fmt.Printf("Length of names array is %v & capacity of names arrays is %v\n", len(names), cap(names))
	// Length of names array is 7 & capacity of names arrays is 7
	// Why is that ? Because arrays have fixed length and capacity equal to their length

	fmt.Printf("\nLength of namesSlice1 is %v & capacity of namesSlice1 is %v\n", len(namesSlice1), cap(namesSlice1))
	// Length of namesSlice1 is 3 & capacity of namesSlice1 is 5
	// Why is that ? Because slices have variable length and capacity equal to how many elements can this slice grow to before it runs out of the underlying array.

	// capacity of a slice basically follows this equation:
	// cap(array[a:b]) = len(array) - a
	// cap(namesSlice1) = len(names) - 2 = 7 - 2 = 5

	fmt.Printf("\nLength of namesSlice3 is %v & capacity of namesSlice3 is %v\n", len(namesSlice3), cap(namesSlice3))
	// Length of namesSlice3 is 4 & capacity of namesSlice3 is 4
	// Why is that ? Because slices have variable length and capacity equal to how many elements can this slice grow to before it runs out of the underlying array.
	// capacity of a slice basically follows this equation:
	// cap(array[a:b]) = len(array) - a
	// cap(namesSlice3) = len(names) - 3 = 7 - 3 = 4

	fmt.Printf("\nLength of namesSlice2 is %v & capacity of namesSlice2 is %v\n", len(namesSlice2), cap(namesSlice2))
	// Length of namesSlice2 is 4 & capacity of namesSlice2 is 7
	// Why is that ? Because slices have variable length and capacity equal to how many elements can this slice grow to before it runs out of the underlying array.
	// capacity of a slice basically follows this equation:
	// cap(array[a:b]) = len(array) - a
	// cap(namesSlice2) = len(names) - 0 = 7 - 0 = 7

	b := [...]int{1, 2, 3, 4, 5} // b is an array of 5 integers
	fmt.Println("dcl:", b)       // dcl: [1 2 3 4 5]

	c := []int{10, 20, 30, 40, 50} // c is a slice of 5 integers
	c = append(c, b[0:]...)        // here we are appending all elements of b to c, we cant directly append b to c because b is an array and c is a slice
	fmt.Println("slc:", c)         // slc: [10 20 30 40 50 1 2 3 4 5]
}

/*
	Limitations (By Design)

	❌ No append
	❌ No delete
	❌ No resize
	❌ No built-in helpers

	Arrays exist mainly to back slices. Rarely used directly.
*/
