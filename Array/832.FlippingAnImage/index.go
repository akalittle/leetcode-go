package Array

func swap(arr *[]int, i, j int) {
	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}

func flipAndInvertImage(arr [][]int) [][]int {
	for _, v := range arr {

		left, right := 0, len(v)-1

		for left < right {
			swap(&v, left, right)
			left++
			right--
		}

		for index := range v {
			v[index] ^= 1
		}

	}

	return arr
}
