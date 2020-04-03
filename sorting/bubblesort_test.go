package sorting

import (
	"sort"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	tests := map[string][]int{
		"PreSorted":  []int{1, 2, 3, 4, 5},
		"Reversed":   []int{5, 4, 3, 2, 1},
		"Inbetween":  []int{3, 5, 1, 2, 4},
		"Dublicates": []int{3, 3, 2, 5, 4},
	}

	//run the tests
	for testcase, list := range tests {
		t.Run(testcase, func(t *testing.T) {
			expected := make([]int, len(list))
			for i, v := range list {
				expected[i] = v
			}

			sort.Ints(expected)

			BubbleSortInt(&list)
			//fmt.Println(testcase, " list after fort: ", list)

			for i := 0; i < len(expected); i++ {
				if list[i] != expected[i] {
					t.Errorf("list[%d] = %d. Expected %d", i, list[i], expected[i])
				}
			}

		})
	}

}
