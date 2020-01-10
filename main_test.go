package main

import "testing"

func Test_findDistance(t *testing.T) {

	tests := []struct {
		name  string
		outgo int
		fuel  int
		want  float64
	}{
		// TODO: Add test cases.
		{"Fuel reserve more than outgo", 10,20,200},
		{"Fuel reserve equal to outgo", 8,8,100},
		{"Fuel reserve less than outgo", 12,6,50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDistance(tt.outgo, tt.fuel); got != tt.want {
				t.Error("findDistance test", tt.name,"want:",tt.want,"got:",got)
			}
		})
	}
}