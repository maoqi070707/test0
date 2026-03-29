package main

import (
	"fmt"
	"math"
	"time"
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

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend!")
	default:
		fmt.Println("It it working day!")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("It is morning")
	case time.Now().Hour() < 18:
		fmt.Println("It is afternoon")
	default:
		fmt.Println("It is evening")
	}

	WhatAnI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Type is int")
		case string:
			fmt.Println("Type is string")
		case bool:
			fmt.Println("Type is bool")
		default:
			fmt.Println("Tyoe is unknown")
		}
	}

	WhatAnI(10)
	WhatAnI("hello")
	WhatAnI(true)

	var arr [5]int
	arr[0] = 1
	arr[1] = 2

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	var twoD = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(twoD)

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			fmt.Println(twoD[i][j])
		}
	}

	slice := make([]string, 5)
	slice[0] = "111"
	slice[1] = "222"
	slice = append(slice, "3", "4", "5")

	slice1 := make([]string, len(slice))
	copy(slice1, slice)
	fmt.Println(slice1)

	s := slice[2:4]
	fmt.Println(s)

	s1 := make([][]int, 3)
	for i := 0; i < len(s1); i++ {
		innerlen := i + 1
		s1[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			s1[i][j] = i + j
		}
	}
	fmt.Println(s1)
}
