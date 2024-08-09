package utils

func ReorderBefore(array []int64, moveID int64, beforeID int64) bool {
	moveIX, beforeIX := 0, 0

	// last element: before out of bounds
	if beforeID == 0 {
		beforeIX = len(array)
	}

	for i := range array {
		if array[i] == moveID {
			moveIX = i
		} else if array[i] == beforeID {
			beforeIX = i
		}
	}

	if moveIX < beforeIX {
		// 1 moveBefore 4 // current = 0 // before = 3
		// old = [1, 2, 3, 4, 5] // new = [2, 3, 1, 4, 5]
		for i := moveIX; i < beforeIX-1; i++ {
			array[i], array[i+1] = array[i+1], array[i]
		}
	} else if moveIX > beforeIX {
		// 4 moveBefore 1 // current = 3 // before = 0
		// old = [1, 2, 3, 4, 5] // new = [2, 3, 1, 4, 5]
		for i := moveIX; i > beforeIX; i-- {
			array[i], array[i-1] = array[i-1], array[i]
		}
	} else {
		return false
	}

	return true
}
