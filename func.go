package main

import "fmt"

func main() {
	num := 6
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by3")
	} else {
		fmt.Println("else")
	}
}
