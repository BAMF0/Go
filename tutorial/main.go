package main

import (
	"fmt"
	"math/rand"
	"math"
)

var c, python, java bool // bools are false by default
var i, j int = 1, 2

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)) // Packages
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(6)) // Imports
	fmt.Println(math.Pi) // Exported names
	fmt.Println(add(42, 13)) // Functions
	a, b := swap("hello", "world") // Multiple results
	fmt.Println(a, b)	
	fmt.Println(split(17)) // Named return values
	var i int // Variables
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!" // Variables with initialisers  
	fmt.Println(i, j, c, python, java) // If an initializer is present, 
	// the type can be omitted; the variable will take the type of the initializer. 
	k := 3 // Short variable declarations: "var" is omitted
	fmt.Println(k) // Outside a function, every statement begins with a keyword,
	// (var, func, and so on) and so the := construct is not available
	

}


// Add takes two parameters of type int and adds them together
func add(x, y int) int { // x int, y int
	return x + y
}

// swap returns two given strings in reverse given order
func swap(x, y string) (string, string) {
	return y, x
}

// split is a short function that displays naked return statements
func split(sum int) (x, y int) {
	x = sum * 4 / 9	
	y = sum - x
	return
}