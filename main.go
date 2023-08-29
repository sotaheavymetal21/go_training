package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println(string("Hello Go"[0]))
	var s string = "Hello Go"

	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "Go"))
	fmt.Println("\"")
	fmt.Println('"')
}
