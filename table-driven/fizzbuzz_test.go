package fizzbuzz_test

import (
	fizzbuzz "github.com/fx2y/hands-on-software-engineering-with-golang/table-driven"
	"testing"
)

func TestFizzBuzzTableDriven(t *testing.T) {
	specs := []struct {
		descr string
		input int
		exp   string
	}{
		{descr: "evenly divisible by 3", input: 9, exp: "Fizz"},
		{descr: "evenly divisible by 5", input: 25, exp: "Buzz"},
		{descr: "evenly divisible by 3 and 5", input: 15, exp: "FizzBuzz"},
		{descr: "edge case", input: 0, exp: "0"},
	}
	for specIndex, spec := range specs {
		if got := fizzbuzz.Evaluate(spec.input); got != spec.exp {
			t.Errorf("[spec %d: %s] expected to get %q; got %q", specIndex, spec.descr, spec.exp, got)
		}
	}
}

func TestFizzBuzzTableDrivenSubtests(t *testing.T) {
	specs := []struct {
		descr string
		input int
		exp   string
	}{
		{descr: "evenly divisible by 3", input: 9, exp: "Fizz"},
		{descr: "evenly divisible by 5", input: 25, exp: "Buzz"},
		{descr: "evenly divisible by 3 and 5", input: 15, exp: "FizzBuzz"},
		{descr: "edge case", input: 0, exp: "0"},
	}
	for specIndex, spec := range specs {
		t.Run(spec.descr, func(t *testing.T) {
			if got := fizzbuzz.Evaluate(spec.input); got != spec.exp {
				t.Errorf("[spec %d: %s] expected to get %q; got %q", specIndex, spec.descr, spec.exp, got)
			}
		})
	}
}
