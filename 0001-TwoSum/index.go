package _001_TwoSum

import "fmt"

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(nums []int, target int) []int {
	store := make(map[int]int, len(nums))

	for i, v := range nums {
		gap := target - v
		if k, ok := store[gap]; ok {
			return []int{k, i}
		}
		store[v] = i
	}
	return nil
}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)

	fmt.Println("res", res)

}
