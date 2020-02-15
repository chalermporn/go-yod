// +build integration

package fizzbuzz_test

import (
	"fmt"
	"hello/fizzbuzz" // name: hello from go.mod,  directory: fizzbuzz
	"testing"
)

func TestFizzBuzz1To100(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Print(fizzbuzz.Say(i) + ", ") // use packet fizzbuzz.go
	}

}

// คำแนะนำ
// ตั้งชืื่อ ตาม packet กับ directory ให้เหมือนกัน
// ชื่อ  directory ควรเป็นตัวเล็กหมด
// func ควรเป็นตัวใหญ่
