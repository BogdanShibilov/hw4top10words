package main

import (
	"reflect"
	"testing"
)

func TestGetTopTenWords(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Empty string",
			args{""},
			[]string{},
		},
		{
			"Multiple words, all have same frequency",
			args{"one two two three THREE tHrEe"},
			[]string{"three", "two", "one"},
		},
		{
			"Multiple words, all have same frequency",
			args{"q w e r t y"},
			[]string{"q", "w", "e", "r", "t", "y"},
		},
		{
			"Single word",
			args{"singular"},
			[]string{"singular"},
		},
		{
			"Cyrillic words",
			args{"один два два три ТРИ тРи"},
			[]string{"три", "два", "один"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTopTenWords(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopTenWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
