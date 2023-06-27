package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var length int = len(arr)
	fmt.Println("Array Length :", length)
	fmt.Println("Array Element :", arr)
}