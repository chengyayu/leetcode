package array

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
	}

	fmt.Println(nums)
}

func Test_MoveZeroes(t *testing.T) {
	arr := []int{1, 2, 0, 0, 5, 6}
	arr1 := []int{0, 0, 0}
	arr3 := []int{1, 2, 5, 6}

	moveZeroes(arr)
	moveZeroes(arr1)
	moveZeroes(arr3)
}
