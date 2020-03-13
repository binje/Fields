package Travel

import "testing"

func TravelVpTest(t *testing.T) {
	expected := []int{0, 0, 1, 3, 3, 3, 4, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10}
	destinations := make([]Destination, 19)

	for i := 0; i < 19; i++ {
		destinations = append(destinations, Hage)
		if TravelVp(destinations) != expected[i] {
			t.Error("Worng Vp travel value")
		}
	}
}
