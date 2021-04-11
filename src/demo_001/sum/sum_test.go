package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("使用 5 个数字的集合", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 6}
		got := Sum(numbers)
		want := 1 + 2 + 3 + 4 + 6
		if want != got {
			t.Errorf("want %d,but got %d! the list is %v", want, got, numbers)
		}
	})

	t.Run("使用 任意 个数字的集合", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 1 + 2 + 3 + 4 + 5 + 6
		if want != got {
			t.Errorf("want %d,but got %d! the list is %v", want, got, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 6}
	got := SumAll(numbers, []int{5, 6, 7})
	want := []int{16, 18}
	// want := "123"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v,but got %v!", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{0})
	want := []int{5, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v,but got %v!", want, got)
	}
}
