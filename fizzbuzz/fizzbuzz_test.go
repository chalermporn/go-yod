package fizzbuzz

import "testing"

func TestFizzBuzzSayOrigin(t *testing.T) {
	t.Run("given 1 say 1", func(t *testing.T) {
		var given = 1
		var want = "1"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 2 say 2", func(t *testing.T) {
		var given = 2
		var want = "2"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
}

func TestFizzBuzzSayFizz(t *testing.T) {
	t.Run("given 3 say Fizz", func(t *testing.T) {
		var given = 3
		var want = "Fizz"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
}
