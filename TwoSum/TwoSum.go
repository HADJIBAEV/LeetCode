package main

import "fmt"

func main() {
	//Example 1
	nums := []int{2, 7, 11, 15}
	target := 9
	//Example 2
	/*nums = []int{3,2,4}
	target := 6 */
	//Example 3
	/*nums := []int{3, 3}
	target := 6 */
	fmt.Println(TwoSum(nums, target))
	fmt.Println(TwoSum2(nums, target))
}

// TwoSum width loop
func TwoSum(nums []int, target int) []int {
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSum2 width map
func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int)     //a+b = target
	for i, num := range nums { //a = nums[i], target   //b = target - a
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}