package algorithms

// SortBubble sorts numbers array in ascending order.
func SortBubble(numbers []int) []int {
	for {
		swapped := false

		for i := 1; i < len(numbers); i++ {
			if numbers[i-1] > numbers[i] {
				numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}

	return numbers
}
