import "slices"

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 1 {
		return 1
	}

	type pair [2]int
	m := make([]pair, len(position))
	for i := range m {
		m[i][0] = position[i]
		m[i][1] = speed[i]
	}
	
	slices.SortFunc(m, func(i, j pair) int {
		if i[0] < j[0] {
			return 1
		} else if i[0] == j[0] && i[1] < j[1] {
			return 1
		} 
		return -1
	})

	// fmt.Println(m)
	cars := make([]pair, 0, len(m))
	for _, v := range m {
		cars = append(cars, v)
	}

	var res int
	var slowestTime float64

	for _, c := range cars {
		iii := float64(target-c[0]) / float64(c[1])   // 10-2 / 5
		if iii > slowestTime {
			res++
			slowestTime = iii
		}
	}

	return res
}