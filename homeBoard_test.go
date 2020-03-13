package main

import "testing"

func TestDikeVp(t *testing.T) {
	for i := 0; i <= 7; i++ {
		if dikeVP(i) != oldDikeVP(i) {
			t.Errorf("Values has changed from %d to %d for input %d", oldDikeVP(i), dikeVP(i), i)
		}
	}
}

func oldDikeVP(i int) int {
	switch i {
	case 7:
		return 0
	case 6:
		return -1
	case 5:
		return -2
	default:
		return -3
	}

}

func TestExtraDikeSpace(t *testing.T) {
	for i := 0; i <= 7; i++ {
		oldV := oldExtraDikeSpace(i)
		newV := extraDikeSpace(i)
		if oldV != newV {
			t.Errorf("Values has changed from %d to %d for input %d", oldV, newV, i)
		}
	}
}

func oldExtraDikeSpace(i int) int {
	switch i {
	case 7:
		return 9
	case 4, 5, 6:
		return 6
	case 1, 2, 3:
		return 3
	default:
		return 0
	}
}
