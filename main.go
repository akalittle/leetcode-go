package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func quick_sort(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	r := rand.Intn(len(nums))

	pivot := nums[r]

	length := len(nums)

	low_part := make([]int, 0, length)
	middle_part := make([]int, 0, length)
	high_part := make([]int, 0, length)

	for _, v := range nums {
		switch {
		case v < pivot:
			low_part = append(low_part, v)
		case v == pivot:
			middle_part = append(middle_part, v)
		case v > pivot:
			high_part = append(high_part, v)
		}
	}

	low_part = quick_sort(low_part)
	high_part = quick_sort(high_part)

	l := make([]int, 0, length)

	l = append(append(append(l, low_part...), middle_part...), high_part...)

	return l

}

func findMinK(nums []int, k int) []int {

	if k >= len(nums) {
		return nums
	}

	return quick_sort(nums)[:k]
}

func i2string(v interface{}) (string, error) {
	switch r := v.(type) {
	case string:
		return r, nil
	case int:
		return strconv.Itoa(r), nil
	}
	return "", nil
}

func main() {
	r, _ := i2string("2333")

	fmt.Println("r", r)
}
