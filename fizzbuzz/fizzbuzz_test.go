package fizzbuzz

import "testing"

func TestFizzBuzzGivenOneSayOne(t *testing.T) {
	var given = 1
	var want = "1"

	var get = Say(given)
	if want != get {
		t.Errorf("given %v want %q but get %q", given, want, get)
	}
}
