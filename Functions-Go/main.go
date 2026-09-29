package main

import (
	"fmt"
	"math"
	"strings"
)
func sayGreetings(name string) {
	fmt.Printf("Hello %s\n", name)
}
func sayBye(name string) {
	fmt.Printf("Goodbye %s\n", name)
}
func cycleNames(name []string) {
	for i := 0; i < len(name); i++ {
		sayGreetings(name[i])
		sayBye(name[i])
	}
}
func cycleNames2(name []string, f func(string)) {
	for _, value := range name {
		f(value)
	}
}
func caclulateCircleArea (radius float64) float64 {
	return math.Pi * radius * radius;
}
func getInitials (s string) (string, string) {
	str := strings.ToUpper(s)
	fmt.Println(str)
	nameSlice := strings.Split(str, " ")
	fmt.Println(nameSlice)
	var initials[] string
	for _, value := range nameSlice {
		initials = append(initials, value[0:1])
	}
	if len(initials) > 1{
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
func main() {
	sayGreetings("Gaurav")
	sayBye("Gaurav")

	names := []string{"Gaurav", "Garima", "Manu", "Tanu"}
	cycleNames(names)
	cycleNames2(names, sayGreetings)
	cycleNames2(names, sayBye)
	radius := 4.58
	fmt.Printf("The area of a circle of radius %v is %0.2f \n",radius, caclulateCircleArea(radius))
	name := "Gaurav Soni"
	initial1, initial2 := getInitials(name)
	fmt.Println(initial1, initial2)
}
