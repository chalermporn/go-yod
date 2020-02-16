package main

import "fmt"

func main() {
	fn := factory()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func factory() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
