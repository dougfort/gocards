package gocards

import (
	"fmt"
	"testing"
)

func TestCard(t *testing.T) {
	testCases := []struct {
		c1          Card
		c2          Card
		expectEqual bool
		expectLess  bool
	}{
		{Card{0, 0}, Card{0, 0}, true, false},
		{Card{0, 1}, Card{0, 0}, false, false},
		{Card{1, 0}, Card{0, 0}, false, false},
		{Card{0, 1}, Card{1, 0}, false, true},
		{Card{1, 1}, Card{1, 10}, false, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("c1 = %s, c2 = %s", tc.c1.Value(), tc.c2.Value()), func(t *testing.T) {
			if tc.c1.Equal(tc.c2) != tc.expectEqual {
				t.Fatalf("Equal failed: %t != %t", tc.c1.Equal(tc.c2), tc.expectEqual)
			}
			if tc.c1.LessThan(tc.c2) != tc.expectLess {
				t.Fatalf("LessThan failed: %t != %t", tc.c1.LessThan(tc.c2), tc.expectLess)
			}
		})
	}
}
