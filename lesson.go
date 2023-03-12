package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	fmt.Println("HelloWorld", time.Now())
	fmt.Println(user.Current())
}

//init() 関数は、パッケージの初期化時に呼び出される関数です。この例では、Init! という文字列が出力されます。
//main() 関数は、アプリケーションのエントリーポイントです。この例では、HelloWorld という文字列と現在の日時が出力されます。また、user.Current() 関数を使用して、現在のユーザー情報を取得し、その情報が出力されます。
