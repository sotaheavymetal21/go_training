package main

// Go言語において、コンパイル時に起動するエントリーポイントは、必ずmainパッケージ内にあるmain()関数です。この関数は、プログラムが起動した際に自動的に呼び出されます。
//この仕組みによって、Go言語では単一のエントリーポイントを持つアプリケーションを作成することができます。また、プログラムがどのように動作するかを明確に示すことができ、可読性が向上するというメリットもあります。

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

var (
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	t   bool    = true
	f   bool    = false
)

func foo() {
	xi := 1
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

const Pi = 3.14

const big = 9223372036854775807 // GOにおけるint型の最大値

// constは型指定をせずGOのコンパイラに解釈はされているが、まだ実行されていない。
//この時点ではoverflowする。printlnするときに-1して実行することでoverflowを防ぐことができる。

const (
	Username = "test_user"
	Password = "test_pass"
)

//【constのメリット】

// 値が変更されないことを保証するため、プログラムが意図しない値に変更されることを防ぐことができる。
// コンパイル時に評価されるため、実行時のパフォーマンスが向上する。
// 名前付き定数を使うことで、可読性が向上し、コードがより分かりやすくなる。
// 【constのデメリット】

// 定数の値が大きい場合、コンパイル時間が長くなることがある。
// 定数を定義するファイルを変更した場合、そのファイルをimportしている全てのファイルが再コンパイルされる。
// 定数を使う場合、定数の値がコンパイル時に評価されるため、実行時に計算する必要がある場合には使えない。

// 【varのメリット】

// 値を変更できるため、動的な値を扱うことができる。
// パッケージ内で定義された変数はデフォルトで初期化されるため、初期値を指定しなくても使うことができる。
// 【varのデメリット】

// 値が変更されることがあるため、意図しない値に変更されることがある。
// 複数のゴルーチンから同時にアクセスされる場合には、競合状態が発生する可能性があるため、注意が必要である。
// 上記のメリット・デメリットを考慮して、constとvarのどちらを使うかは、その変数や定数がどのような役割を持ち、どのように使われるかによって異なります。定数であることが望ましい場合にはconstを、動的な値を扱う必要がある場合にはvarを使うことが適切であると言えます。

func main() {
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
	fmt.Println(i, f64, s, t, f)
	foo()
	fmt.Println(Pi, Username, Password)
	fmt.Println(big)
}
