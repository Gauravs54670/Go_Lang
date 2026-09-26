package main

import (
	"fmt"
	"sort"
	"strings"
)
func main() {
	// strings are immutable like in Java so these functions always returns a new string
	message := "Hello, This is Gaurav"
	fmt.Println(message)
	fmt.Println(strings.Contains(message, "hello"))
	fmt.Println(strings.Contains(message, "Hello"))
	str := "Go is good, Go is fast, Go is good programming language."
	// Replaces() replace a specified number of occurrences
	// 1 -> replace only first occurrences
	fmt.Println(strings.Replace(str,"Go","Java", 1))
	fmt.Println(strings.Replace(str, "Go", "Java", 2))
	fmt.Println(strings.Replace(str,"Go", "Java", 3))
	// ReplacesAll() replaces all occurrences
	fmt.Println(strings.ReplaceAll(str, "Java", "Go"))
	fmt.Println(strings.ToUpper(str) +" "+strings.ToLower(str))
	fmt.Println(strings.Index(str, "G"))
	fmt.Println(strings.Index(str, "Go"))
	fmt.Println(strings.Index(str, "go"))
	fmt.Println(strings.Index(str, "good"))
	fmt.Println(strings.Index(str, "programming"))
	// strings.Split() splits a string into a slice of strings based on the given separator.
	fmt.Println(strings.Split(str, " "))
	fmt.Println(strings.Split(str, ","))

	numbers := []int{87,98,24,5,47,96,45}
	// sort in asscending order
	sort.Ints(numbers)
	fmt.Println(numbers)
	names := []string{"Gaurav", "Amit", "Rahul"}
	// sort strings alphabetically
	sort.Strings(names)
	fmt.Println(names)
	// Binary search on a sorted slice.
	// Returns the index if found, otherwise the position where it can be inserted.
	index := sort.SearchInts(numbers, 24)
	fmt.Println(numbers[index])
	index = sort.SearchInts(numbers, 100)
	fmt.Println(len(numbers), "is the length of a numbers slice and because there is no 100 present it will return the index where 100 can be inserted.", index)
	index = sort.SearchStrings(names, "Rahul")
	fmt.Printf("Index is %v Values is %s\n",index, names[index])
}