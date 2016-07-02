package main

import "testing"


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
				t.Errorf("Reverse(%q) == %q, want %q", c.israndom, c.length, c.currentindex, c.want)
			}
		}
}