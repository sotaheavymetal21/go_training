package main

import "fmt"

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

// intのマックス数値、オーバーフローする
// var big int = 9223372036854775807 + 1
// 解釈はされているが、コンパイラはされていない状態
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
