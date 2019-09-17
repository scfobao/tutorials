package main

import "fmt"

func main() {
	fmt.Println("test")
	var s string = "我是一个中国人"
	var i int = 30
	var f int = 638
	fmt.Println("fmt print:", s)
	fmt.Println(i, f)
	k := i + f
	fmt.Print(k)
	var m, n int = 1, 2
	var s1, s2, s3 = "I", " am ", "liqinghua"
	fmt.Println(m, n)
	fmt.Println(s1, s2, s3)
}
