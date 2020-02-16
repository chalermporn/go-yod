package fizzbuzz

import "strconv"

// type FizzBuzzOop struct {
// 	Number int `json:"number"`,
// }
// func (n FizzBuzzOop) New() int {
// 	return n.Number
// }

type FizzBuzz int

func New(n int) FizzBuzz {
	return FizzBuzz(n)
}
func (f FizzBuzz) String() string {
	return Say(int(f))
}

func Say(n int) string {
	if n%15 == 0 { // seem n%15
		return "FizzBuzz"
	}
	// if n%3 == 0 && n%5 == 0 { // seem n%15
	// 	return "FizzBuzz"
	// }
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

// 3 f
// 5 b
// 3 5 fb
