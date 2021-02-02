package main

import "fmt"

func main() {
	fmt.Println("Enter the number:\n")

	var number int
	fmt.Scanln(&number)

	fmt.Println("The converted form of the number in decimal,hexa and binary is\n")

	fmt.Printf("%d\n,%x\n,%b\n", number, number, number)
}
