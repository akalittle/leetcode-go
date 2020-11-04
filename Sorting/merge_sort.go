package Sorting

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(l, r []int) []int {
	result := make([]int, 0, len(l)+len(r))

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			result = append(result, r...)
			return result
		}
		if len(r) == 0 {
			result = append(result, l...)
			return result
		}
		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	return result
}
