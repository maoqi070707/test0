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

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 3; i++ {
		fmt.Print("===")
	}

	for {
		println("loop")
		break
	}

	for i := 3; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("=-=-=-=-=-=-=-")
		}
	}

	if num := 10; num > 5 {
		fmt.Println("bigger than 5")
	} else if num > 9 {
		fmt.Println("bigger than 9")
	} else {
		fmt.Println("smaller than 5 or 9")
	}

}
