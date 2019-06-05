package twosum

import "testing"
import . "../common"

func TestTwoSum(t *testing.T) {
	tables := []struct {
		nums   []int
		target int
		result []int
	}{
		// {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, table := range tables {
		result := twoSum(table.nums, table.target)
		if !Equal(result, table.result) {
			t.Errorf("the result %v in target %d to find in nums %v is not equal to true answer %v", result, table.target, table.nums, table.result)
		}
	}
}
