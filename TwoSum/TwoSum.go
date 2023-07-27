package main

import "fmt"

func main() {
	nums := [...]int{2, 7, 11, 15}
	target := 9
	fmt.Println(TwoSum(nums, target))
}
func TwoSum(nums [4]int, target int) []int {
	m := make(map[int]int)     //a+b = target
	for i, num := range nums { //a = nums[i], target   //b = target - a
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
