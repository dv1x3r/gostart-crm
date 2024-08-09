package utils

import "testing"

func TestReorderBefore(t *testing.T) {
	tests := []struct {
		InputArray     []int64
		InputCurrentID int64
		InputBeforeID  int64
		ExpectedArray  []int64
		ExpectedOk     bool
	}{
		{
			InputArray:     []int64{1, 2, 3, 4, 5},
			InputCurrentID: 1,
			InputBeforeID:  0,
			ExpectedArray:  []int64{2, 3, 4, 5, 1},
			ExpectedOk:     true,
		},
		{
			InputArray:     []int64{1, 2, 3, 4, 5},
			InputCurrentID: 5,
			InputBeforeID:  1,
			ExpectedArray:  []int64{5, 1, 2, 3, 4},
			ExpectedOk:     true,
		},
		{
			InputArray:     []int64{1, 2, 3, 4, 5},
			InputCurrentID: 2,
			InputBeforeID:  5,
			ExpectedArray:  []int64{1, 3, 4, 2, 5},
			ExpectedOk:     true,
		},
		{
			InputArray:     []int64{1, 2, 3, 4, 5},
			InputCurrentID: 5,
			InputBeforeID:  2,
			ExpectedArray:  []int64{1, 5, 2, 3, 4},
			ExpectedOk:     true,
		},
		{
			InputArray:     []int64{1, 2, 3, 4, 5},
			InputCurrentID: 1,
			InputBeforeID:  1,
			ExpectedArray:  []int64{1, 2, 3, 4, 5},
			ExpectedOk:     false,
		},
		{
			InputArray:     []int64{1, 2, 3, 4, 5},
			InputCurrentID: 1,
			InputBeforeID:  10,
			ExpectedArray:  []int64{1, 2, 3, 4, 5},
			ExpectedOk:     false,
		},
		{
			InputArray:     []int64{1, 2, 3, 4, 5},
			InputCurrentID: 1,
			InputBeforeID:  -1,
			ExpectedArray:  []int64{1, 2, 3, 4, 5},
			ExpectedOk:     false,
		},
	}

	for ti, test := range tests {
		ok := ReorderBefore(test.InputArray, test.InputCurrentID, test.InputBeforeID)
		if ok != test.ExpectedOk {
			t.Fatalf("[%d] ok does not match, actual [%+v], expected [%+v]\n", ti, ok, test.ExpectedOk)
		}

		if len(test.InputArray) != len(test.ExpectedArray) {
			t.Fatalf("[%d] len does not match, actual [%+v], expected [%+v]\n", ti, len(test.InputArray), len(test.ExpectedArray))
		}

		for i := range test.InputArray {
			if test.InputArray[i] != test.ExpectedArray[i] {
				t.Fatalf("[%d] value does not match, actual [%+v], expected [%+v]\n", ti, test.InputArray, test.ExpectedArray)
			}
		}
	}
}
