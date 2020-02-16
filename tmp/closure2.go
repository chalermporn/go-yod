package main

import "fmt"

func main() {
	fn1, fn2 := factory()
	fn1()
	fn1()
	fmt.Println(fn2())

	fn1()
	fmt.Println(fn2())
}

func factory() (func(), func() int) {
	var i int
	return func() {
			i++
		},
		func() int {
			return i
		}
} // END
