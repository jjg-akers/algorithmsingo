package sorting

//InsertionSortInt will use insertion sort on a list of ints
// O(N^2)
func InsertionSortInt(list []int) {
	var sortedList []int
	for _, val := range list {
		for j, val2 := range sortedList {
			if val < sortedList[val2] {
				sortedList = append(sortedList[:j], append([]int{val}, sortedList[j:]...)...)
			}
		}
		sortedList = append(sortedList, val)
	}

	// copy back to original list
	for i, val := range sortedList {
		list[i] = val
	}
}
