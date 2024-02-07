package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func IsPalindrome(str string) bool {
	lastIndex := len(str) - 1
	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[lastIndex-i] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	t.Run("scenario1", func(t *testing.T) {
		input := "radar"
		want := true

		have := IsPalindrome(input)
		if have != want {
			t.Errorf("expected %t get %t", want, have)
		}
	})
}

func FindMax(in []int) int {
	sort.Ints(in)
	return in[len(in)-1]
}

func TestIsMax(t *testing.T) {
	t.Run("scenario1", func(t *testing.T) {
		input := []int{3, 5, 1, 9, 2}
		want := 9

		have := FindMax(input)
		if have != want {
			t.Errorf("expected %d get %d", want, have)
		}
	})
}

func PrintTriangle(in int) {
	for i := 1; i <= in; i++ {
		fmt.Printf("%s\n", strings.Repeat("*", i))
	}
}

func TestPrintTriangle(t *testing.T) {
	t.Run("scenario1", func(t *testing.T) {
		input := 5

		PrintTriangle(input)
	})
}

func Factorial(in int) int {
	if in < 1 {
		return 1
	}

	return in * Factorial(in-1)
}

func TestFactorial(t *testing.T) {
	t.Run("scenario1", func(t *testing.T) {
		input := 5
		want := 120

		have := Factorial(input)
		if have != want {
			t.Errorf("expected %d get %d", want, have)
		}
	})
}
