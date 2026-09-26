package main

import "fmt"

func main() {
	var ages[3]int = [3]int{25,30,75}
	fmt.Println(ages)
	var counters[5]int = [5]int{1,2,3,4,5}
	fmt.Println(counters)
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(numbers)
	fmt.Println(numbers, len(ages))
	names := []string{"Gaurav", "Garima"}
	fmt.Println(names, len(names))
	fmt.Printf("My name is: %s\n and my sister name is: %s\n",names[0], names[1])
	fmt.Println(append(numbers, 58))
	fmt.Println(append(names, "Riya"))
	// for loop
	rangeOne := counters[0:5]
	fmt.Println(rangeOne)
	rangeOne = counters[1:4]
	fmt.Println(rangeOne)
	rangeTwo := counters[:]
	fmt.Println(rangeTwo)
	rangeThree := counters[1:]
	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Println(len(rangeOne), cap(rangeOne))
	fmt.Println(append(rangeOne, 58))
	// cap = 4
	/* 
		var counters[5]int = [5]int{1,2,3,4,5}
		rangeOne = counters[1:4]
		why cap = 4 ? because rangeOne slice starts from 1st index and 
		From index 1 to the end, there are 4 positions available
	*/
	fmt.Println(len(rangeOne), cap(rangeOne))
	/*
		append() does not modify the length of the original slice variable 
		unlike other Programming Language.
		that's why rangeOne still has 3 length counters[1:4]
	*/
	fmt.Println(rangeOne)
	// .append() does not append into the original slice instead it returns the updated new slice
	rangeFour := append(rangeOne, 78)
	fmt.Println(rangeFour)
	fmt.Println(len(rangeFour))
}