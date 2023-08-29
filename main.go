package main

import (
	"fmt"
	"time"
)

// i5 := 500
var i5 int = 100

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)

	var str string = "Hello Go"
	fmt.Println(str)

	var (
		i2 int
		s2 string
	)
	i2 = 300
	s2 = "Go"
	fmt.Println(i2, s2)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義、型推論
	// 変数名 := 値

	i4 := 400
	fmt.Println(i4)

	fmt.Println(i5)

	outer()

}
