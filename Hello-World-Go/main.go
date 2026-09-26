package main
// for formating strings and printing messages to the console
import "fmt"

func main() {
	fmt.Println("Hello World!!!")
	var variableName1 string = "Go string variable name 1"
	var variableNam2 = "Go string variable name 2"
	fmt.Println(variableName1, variableNam2)
	// fmt.Printf call needs 1 arg but has 2 args
	/*fmt.Printf("1st variable name is %s", variableName1+"\n" +
					"Second variable name is %s", variableNam2)
					This is not allowed with Printf()
					*/
	fmt.Printf("First variable name is %s\nSecond variable name is%s\n",variableName1,variableNam2)
	// same as var string variableName 3
	varialbeName3 := "Go string variable 3"
	fmt.Println(varialbeName3)
	varialbeName3 = "Go string variable 3 assigned"
	fmt.Print(varialbeName3+"\n")
	var number1 int = 1
	var number2 = 2
	number3 := 3
	fmt.Println(number1, number2, number3)
	// bit size and memory
	// Go doesn't use long type of variable instead it uses the bits so each sections of bits 
	// i.e; 8 bits, 16 bits to store the long numbers
	// var number4 int8 = 25554 -> will throw an error 
	var number4 int16 = 25554
	fmt.Println(number4)
	number5 := 87665548
	fmt.Println(number5)
	// var number6 uint16 = -55655 -> throw error uint can't store negative numbers
	var number6 uint = 478
	fmt.Println(number6)
	number7 := -785
	fmt.Println(number7)
	// when declaring float declaring the bit size is necessary
	var number8 float32 = 78.9
	var number9 float32 = -78.9
	number10 := 795.7
	fmt.Println(number8, number9, number10)
	name := "Gaurav"
	age := 22
	fmt.Println("My name is ",name, "My age is ",age)
	fmt.Printf("My name is %v and My age is %v \n", name, age)
	// %s is mainly for string type
	fmt.Printf("My name is %s and My age is %v \n", name, age)
	// %q puts the string only into the quots ("")
	fmt.Printf("My name is %q and My age is %v \n", name, age)
	// %T defins the variable type of a variable
	fmt.Printf("Variable type of age is %T Variable type of name is %T\n", age, name)
	// %f for float types
	distance := 78.5
	fmt.Printf("Distance is %f\n", distance)
	fmt.Printf("Distance is %0.2f\n", distance)

	// Sprintf -> save formated string and return the formated string
	// does not print anything just return the stored string
	str := fmt.Sprintf("My name is %s and My age is %v\n", name, age)
	fmt.Println("Stored string is: ", str)
}