package main

import (
	"fmt"
	"math"
	"time"
	"unicode/utf8"
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

	//:= 只能在函数内使用（局部变量）
	//var全局或局部均可
	//:= 允许在同一作用域「部分新变量」
	//var 同一作用域不能重复
	//:= 必须马上赋值
	//var可声明不赋值
	//多返回值场景:=更优雅
	num := []int{1, 2, 3, 4, 5}
	for index, value := range num {
		fmt.Println(index, value)
	}

	dict := make(map[string]int)
	dict["a"] = 1
	dict["b"] = 2
	dict["c"] = 3

	for key := range dict {
		fmt.Println("key:", key)
	}

	for key, value := range dict {
		fmt.Printf("%s -> %d\n", key, value)
	}

	str := "hello你好😊"
	for i, c := range str {
		fmt.Printf("%d %c\n", i, c)
	}

	result := plus(1, 2)
	fmt.Println(result)
	fmt.Println(plusplus(1, 2, 3))
	fmt.Println(swap("hello", "world"))
	n1, _ := swap("hello", "world")
	fmt.Println(n1)

	ns := []int{1, 2, 3, 4, 5}
	ns1 := append(ns, 6, 7, 8)
	fmt.Println(sum(ns...))
	fmt.Println(sum(ns1...))

	//闭包 = 函数 + 它引用的外部环境变量
	//闭包会记住变量，不会被垃圾回收
	//闭包可以封装私有状态
	//闭包引用的是变量本身，不是值（注意坑）
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("------------------------")

	var functions []func()
	//Go 1.22 版本对 for 循环做了重大优化：
	//循环变量 i 每次迭代都会创建一个新的副本，而不是复用同一个变量。
	//因此，在循环体中对变量 i 的修改，不会影响到外层循环的下一次迭代。
	for i := 0; i < 3; i++ {
		functions = append(functions, func() {
			fmt.Println(i)
		})
	}
	for _, f := range functions {
		f()
	}

	fmt.Println(fact(5))
	var fib func(w int) int
	fib = func(w int) int {
		if w < 2 {
			return w
		}
		return fib(w-1) + fib(w-2)
	}

	fmt.Println(fib(7))

	I := 10
	Zeroval(&I)
	fmt.Println(I)

	str1 := "สวัสดี"
	fmt.Println("rune count:", utf8.RuneCountInString(str1))
	//range函数遍历字符串
	for i, r := range str1 {
		fmt.Printf("%d -> %c\n", i, r)
	}
	//手动解码
	for i, w := 0, 0; i < len(str1); i += w {
		runeValue, width := utf8.DecodeRuneInString(str1[i:])
		fmt.Printf("%d -> %c\n", i, runeValue)
		w = width
		examineRune(runeValue)
	}

	fmt.Println(person{"Tom", 25})
	fmt.Println(person{name: "Jerry", age: 30})
	fmt.Println(person{name: "Bob"})
	fmt.Println(&person{name: "Alice", age: 20})
	sss := newPerson("Sam")
	fmt.Println(sss)
	fmt.Println(sss.name)

	r := rect{3.0, 4.0}
	fmt.Println(r.area())
	r.setArea(10.0)
	fmt.Println(r.wide, r.high)

	measure(r)
	measure(&r)
}

func newPerson(name string) *person {
	p := new(person)
	p.name = name
	p.age = 23
	return p
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (r rect) perim() float64 {
	return 2*r.wide + 2*r.high
}

type rect struct {
	wide, high float64
}

func (r rect) area() float64 {
	return r.wide * r.high
}

func (r *rect) setArea(area float64) {
	r.wide = math.Sqrt(area)
	r.high = math.Sqrt(area)
}

type person struct {
	name string
	age  int
}

func examineRune(r rune) {
	if r == 'ส' {
		fmt.Println("found")
	}
}
func Zeroval(ptr *int) {
	*ptr = 0
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

// 多返回值
func swap(x, y string) (string, string) {
	return y, x
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
