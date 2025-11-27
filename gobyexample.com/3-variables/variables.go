package main


import "fmt"

func main(){
	var a, b int = 3, 5
	var c = "hi there"
	var d = true
	var e float64 = 3.14

	fmt.Println(a, b, c, d, e)

	f := "shorthand for declaring and initializing a variable"
	fmt.Println(f)

	g, h := 3, 4
	fmt.Println(g, h)
}