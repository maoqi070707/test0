package main

import (
	"fmt"
	"math"
)

const s = "content"

func main() {
	fmt.Println("hello world")
	fmt.Println("Go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println(!true)
	fmt.Println("7.0/3.0", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(true && true || false)

	var a int = 10
	fmt.Println(a)

	var b, c int = 10, 20
	fmt.Println(b + c)

	var d float32 = 3.14
	fmt.Println(d)

	var e bool = true
	fmt.Println(e)

	var f string = "hello"
	fmt.Println(f)
	//简写
	g := "world"
	fmt.Println(f + " " + g)

	const n = 500000000
	const m = 6e20 / n
	fmt.Println(m)
	fmt.Println(int64(m))

	fmt.Println(math.Pi)
	fmt.Println(math.Sin(n))

	fmt.Println(s)
}
