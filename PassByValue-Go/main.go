package main

import (
	"fmt"
)
// a new memory location is assigned for this s variable so 
// the changess on this variable will not be reflected on the main s variable 
func passByValue(s string) {
	s = "New Value"
}
// Map is passed by value, but a copy of its pointer/reference to the
// actual map data is passed. So both point to the same underlying data,
// which is why changes made inside a function affect the original map.
//
// Unlike basic/value types like int, float, and bool, map is a
// reference-like type in Go.
func passByValueMaps(menu_and_price map[string]float32) {
	menu_and_price["Chicken"] = 600.450
	menu_and_price["Dosa"] = 453.45
}
func main() {

	fmt.Println("Hello World!!!")

	s := "Hello Gaurav"
	// A copy of s variable is created and a copy is passed not the original string
	passByValue(s)
	fmt.Println(s)

	menu_and_price := map[string]float32 {
		"Rice": 124.50,
		"Curry": 230.45,
		"Chicken": 450.45,
	}

	fmt.Println(menu_and_price)

	
	passByValueMaps(menu_and_price)

	fmt.Println(menu_and_price)

}