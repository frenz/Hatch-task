package main

import (
	"testing"
)

func Test_compareHashMaps(t *testing.T) {
	type args struct {
		mapSource map[string]bool
		mapTarget map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"f", args{map[string]bool{}, map[string]bool{"foo": true}}, false},
		{"foo", args{map[string]bool{"foo": true}, map[string]bool{"foo": true}}, true},
		{"foo bar", args{map[string]bool{"foo": true}, map[string]bool{"bar": true}}, false},
		{"foo bar foo", args{map[string]bool{"foo": true, "bar": true}, map[string]bool{"foo": true}}, false},
		{"foo bar bar foo", args{map[string]bool{"foo": true, "bar": true}, map[string]bool{"bar": true, "foo": true}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareHashMaps(tt.args.mapSource, tt.args.mapTarget); got != tt.want {
				t.Errorf("compareHashMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
