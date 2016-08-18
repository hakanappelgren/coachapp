package main

import "testing"

//Only test for static index, not random
// @ToDo test case for random set-ut
func TestGetIndex (t *testing.T) {
	cases := []struct {
		israndom bool;
		length int;
		currentindex int;
		want int
	}{
		{false, 3, 0, 1},
		{false, 3, 1, 2},
		{false, 3, 2, 0},
	}
		for _, c := range cases {
			got := GetIndex(c.length, c.israndom, c.currentindex)
			if got != c.want {
				t.Errorf("Israndom: (%q) Length: %q Got: %q Want: %q", c.israndom, c.length, got, c.want)
			}
		}
}

func TestGetQuestion (t *testing.T) {
	cases := []struct {
		index int;
		israndom bool;
		want string
	}{
		{0, false, "andra frågan"},
		{1, false, "tredje frågan"},
		{2, false, "första frågan"},
		{0, true, "andra frågan"},
		{1, true, "andra frågan"},
		{2, true, "tredje frågan"},
	}
	for _, c := range cases {
		got := GetQuestion (c.israndom, c.index)
		if got != c.want {
			t.Errorf("Function: GetQuestion: got %q for index %q, want %q", got, c.index, c.want)
		}
	}
}
