package main

import "fmt"

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
func main() {
	sayGreetings("Gaurav")
	sayBye("Gaurav")

	names := []string{"Gaurav", "Garima", "Manu", "Tanu"}
	cycleNames(names)
	cycleNames2(names, sayGreetings)
	cycleNames2(names, sayBye)
}
