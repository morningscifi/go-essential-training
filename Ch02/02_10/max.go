// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println(nums)
	max_index := find_largest(nums)
	fmt.Printf("Nums[%d]= %d", max_index, nums[max_index])
}

func find_largest(nums []int) int {
	idx := 0
	for i, val := range nums {
		if val >= nums[idx] {
			idx = i
		}
	}
	return idx
}
