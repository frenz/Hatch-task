package main

import (
	"reflect"
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

func Test_streamToMapHash(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantRes map[string]bool
		wantErr bool
	}{
		{"foo", args{"./data/input2.json"}, map[string]bool{"65656636323230396538393737393039313462303762623562383766363132373939383662383864": true, "64383166343664386465653835613333343261396335313830356464643339643939393865386236": true}, false},
		{"bar", args{"./data/input1.json"}, map[string]bool{"65656636323230396538393737393039313462303762623562383766363132373939383662383864": true, "64383166343664386465653835613333343261396335313830356464643339643939393865386236": true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := streamToMapHash(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("streamToMapHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("streamToMapHash() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
