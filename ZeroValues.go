// go is all about and oriented about the type

//declaring values of int

package main

import "fmt"

var y = 42
var string1 = "tttt"

// string1 = "tttt" - statically typed language, unlikely loosely typed languages

//declare a variable of certain type and then ASSIGN a value of that type to the variable
var abc string
var a int

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) //type
	fmt.Println(string1)
	fmt.Printf("%T\n", string1)
	fmt.Println(abc)
	fmt.Println(a)
}
