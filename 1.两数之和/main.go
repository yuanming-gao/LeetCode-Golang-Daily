package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {

		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}

		m[v] = i
	}
	return nil
}
