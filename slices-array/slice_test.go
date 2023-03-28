package listpractice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checksums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v, got: %v", want, got)
		}

	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection with multiple arrays", func(t *testing.T) {
		want := SumAll([]int{1, 2}, []int{2, 3})
		got := []int{3, 5}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("checking 1: values inide the objects", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{3, 4})
		want := []int{5, 4}

		checksums(t, got, want)
	})

}
