func twoSum(numbers []int, target int) []int {
	ptr1, ptr2 := 0, len(numbers)-1

	for ptr1 < ptr2 {
		if numbers[ptr1]+numbers[ptr2] == target {
			return []int{ptr1+1, ptr2+1}
		} else if numbers[ptr1]+numbers[ptr2] < target {
			ptr1++
		} else {
			ptr2--
		}
	} 
	
	return []int{-1, -1}
}
