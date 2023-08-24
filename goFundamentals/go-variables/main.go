package main

import "fmt"

// d := 5 Contoh pemakaian variable yang salah
var f = "contohnya"
var m string

func main() {
	var a int = 10
	var b = "saya"
	c := 88

	var g,h,i,j = 1, "saya", "bwaa", false

	var (
		k string = "ini di kurungin"
		l float64 = 0.15
	)

	m = "Variable Luar"

	var n string
	var o int
	var p bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(f)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println(k)
	fmt.Println(l)

	fmt.Println(m)

	// isi default dari macam macam variable tipe data
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
}