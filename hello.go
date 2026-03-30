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

	i1 := []int{1, 2, 3, 4, 5}
	i2 := i1[1:4]
	fmt.Println(i2)
	fmt.Println("切片长度不固定")

	i2 = make([]int, len(i1), 10)
	fmt.Println("切片长度固定,第三个参数为容量")
	fmt.Println(len(i2), cap(i2))

	// 切片的容量不足时，会自动扩容，扩容的大小为原来的2倍
	x := make([]int, 3, 5)
	t := make([]int, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	for i := range s {
		t[i] = x[i]
	} //for循环可用copy简化
	x = t

	//向切片的末尾加入数据，如果容量不足，会自动扩容
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)

	//创建空map需要用make
	mp := make(map[string]int)

	//创建非空map
	mp1 := map[string]int{"1st": 1, "2nd": 2}
	fmt.Println(mp1)
	mp["1st"] = 1
	mp["2nd"] = 2
	fmt.Println(mp)

	//判断键是否存在
	_, prs := mp["3rd"]
	fmt.Println(prs)

	delete(mp, "2nd")

	//访问时，要有判断键是否存在
	val, prs := mp["1st"]
	fmt.Println(val, prs)
	fmt.Println("len(mp)=", len(mp))
}

// AppendByte 向切片的末尾加入数据，如果容量不足，会自动扩容
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
