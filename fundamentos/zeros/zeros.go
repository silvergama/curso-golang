package main

import "fmt"

func main() {
	var a int // quando não é atribuido um valor é necessário informar o tipo
	var b float64
	var c bool
	var d string
	var e *int
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e) // 0 0 false "" <nil>
}
