package main

import "fmt"

const (
	PI        = 3.1415926
	BATH_ROOT = "/etc/nginx/conf"
	APPNAME   = "CONS"
)

func main() {
	fmt.Println(PI)
	fmt.Println(BATH_ROOT)
	fmt.Println(APPNAME)
	const a = 12.34
	fmt.Println(a)
	fmt.Println(float32(a))
	hello()
}

func hello() {
	fmt.Println(PI)
	fmt.Println(BATH_ROOT)
	fmt.Println(APPNAME)
	//fmt.Println(a)
}
