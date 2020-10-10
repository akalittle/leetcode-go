package Array

import (
	"sort"
)

//Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
//Each number in candidates may only be used once in the combination.
//
//Note:
//
//All numbers (including target) will be positive integers.
//The solution set must not contain duplicate combinations.
//Example 1:
//
//Input: candidates =  [10, 1,2, 7, 6, 1, 5], target =  8,
//A solution set is:
//[
//[1, 7],
//[1, 2, 5],
//[2, 6],
//[1, 1, 6]
//]
//Example 2:
//
//Input: candidates =  [2, 5, 2, 1, 2], target =  5,
//A solution set is:
//[
//  [1, 2, 2],
//  [5]
//]

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	comb, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findCombinationSum2(candidates, target, 0, comb, &res)
	return res
}

func findCombinationSum2(candidates []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			// 本次不取重复数字，下次循环可能会取重复数字
			continue
		}
		if target >= candidates[i] {
			c = append(c, candidates[i])
			findCombinationSum2(candidates, target-candidates[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
