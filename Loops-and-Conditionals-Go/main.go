package main

import "fmt"

func main() {
	// while
	x := 0
	for x < 5 {
		fmt.Println("Value of x is ", x)
		x++
	}
	fmt.Println()
	// for
	for i := 1; i<=5; i++ {
		fmt.Println("Value of i is ", i)
	}
	names := []string{"Gaurav", "Garima", "Anshika", "Richa", "Pihu"}
	for i := 0; i<len(names); i++ {
		fmt.Println(names[i])
	}
	
	for index, value := range names {
		fmt.Printf("The index is %v and the value is %v\n", index, value)
	}
	if x == 5 {
		fmt.Printf("%v is equalls to 5 \n",x)
	} else if x < 5 {
		fmt.Printf("%v is less than 5\n",x)
	} else {
		fmt.Println("Good Bye")
	}
	numbers := []int{45,89, 10, 15, 99, 2, 4, 5, 20}
	for i:= 0; i< len(numbers); i++ {
		if numbers[i] < 20 && numbers[i] != 5{
			fmt.Printf("Continuing the index %v where value is %v\n", i, numbers[i])
			continue
		} else if numbers[i] == 5 {
			fmt.Printf("Breaking the loop at index %v where value is %v\n", i, numbers[i])
			break
		} else {
			fmt.Println(numbers[i])
		}
	}
}