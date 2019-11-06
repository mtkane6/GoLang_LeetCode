package solutions

// MergeSort implementation in Go
func MergeSort(a []int) []int {
	// first split slice recursively
	length := len(a)
	if length == 1 {
		return a
	}
	middle := int(length / 2)

	leftArray := (make([]int, middle))
	rightArray := (make([]int, length-middle))

	// separate current array into left and right portions
	for i := range a {
		if i < middle {
			leftArray[i] = a[i]
		} else {
			rightArray[i-middle] = a[i]
		}
	}

	return merge(MergeSort(leftArray), MergeSort(rightArray))
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))

	// sort the Left and Right arrays into result array
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// add the rest of Left to the result array
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	// add the rest of Right to the result array
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
