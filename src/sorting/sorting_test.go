package sorting

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	nums := []int{4, 8, 10, -1, 100}
	expectedNums := []int{-1, 4, 8, 10, 100}
	BubbleSort(nums)
	if !reflect.DeepEqual(nums, expectedNums) {
		t.Error("BubbleSort failed to sort nums", nums)
	}
}
