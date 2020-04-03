package sorting

// BubbleSortInt will sort a list of ints using... Bublle Sort
//O(N^2)
// good for small lists, and nearly sorted lists
func BubbleSortInt(list *[]int) {
	for i := 0; i < len(*list); i++ {
		swap := false
		for j := 0; j < len(*list)-1-i; j++ {
			if (*list)[j] > (*list)[j+1] {
				// swap
				(*list)[j], (*list)[j+1] = (*list)[j+1], (*list)[j]
				swap = true
			}
			//fmt.Println(list)
		}
		if !swap {
			break
		}
	}

}

func cmp(first, second int) int {
	if first > second {
		return 1
	} else if first < second {
		return -1
	} else {
		return 0
	}
}
