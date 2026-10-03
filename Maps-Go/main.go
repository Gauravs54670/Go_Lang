package main

import (
	"fmt"
)

func main() {
	menu_price := map[string]float32 {
		"Rice": 42.78,
		"Curray": 78.5,
		"Soba": 120.5,
		"Sushi": 78.6,
	}
	fmt.Println(menu_price)
	for key, value := range menu_price {
		fmt.Println(key, " : ", value)
	}
	class := map[int]string {
		1 : "Gaurav Soni",
		2 : "Garima Soni",
		3 : "Manu Soni",
		4 : "Tanu Soni",
	}
	fmt.Println("Roll Number : ", "Name")
	for key, value := range class {
		fmt.Print(key, "	", value, "\n")
	}
	fmt.Println(class[1])
}