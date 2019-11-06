package solutions

// Quicksort is the Go implementation of Quicksort
func Quicksort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	leftIndex, rightIndex := 0, len(a)-1

	pivotPoint := len(a) / 2

	a[pivotPoint], a[rightIndex] = a[rightIndex], a[pivotPoint]

	for i := range a {
		if a[i] < a[rightIndex] {
			a[leftIndex], a[i] = a[i], a[leftIndex]
			leftIndex++
		}
	}
	a[leftIndex], a[rightIndex] = a[rightIndex], a[leftIndex]

	Quicksort(a[:leftIndex])
	Quicksort(a[leftIndex+1:])

	return a
}
