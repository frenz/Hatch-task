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

func Test_readBytes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRes map[string]bool
	}{
		{"bar", args{[]byte("{\"name\":\"test json\",\"id\":\"jhasdad\"}")}, map[string]bool{"65656636323230396538393737393039313462303762623562383766363132373939383662383864": true}},
		{"foo", args{[]byte("{\"id\":\"wqweq\",\"name\":\"test json 2\"}")}, map[string]bool{"64383166343664386465653835613333343261396335313830356464643339643939393865386236": true}},
		{"foobar", args{[]byte("[{\"id\":\"wqweq\",\"name\":\"test json 2\"},{\"name\":\"test json\",\"id\":\"jhasdad\"}]")}, map[string]bool{"65656636323230396538393737393039313462303762623562383766363132373939383662383864": true, "64383166343664386465653835613333343261396335313830356464643339643939393865386236": true}},
		{"barfoo", args{[]byte("[{\"name\":\"test json\",\"id\":\"jhasdad\"},{\"id\":\"wqweq\",\"name\":\"test json 2\"}]")}, map[string]bool{"65656636323230396538393737393039313462303762623562383766363132373939383662383864": true, "64383166343664386465653835613333343261396335313830356464643339643939393865386236": true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := readBytes(tt.args.data); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("readBytes() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_fileJSONToByte(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"foo", args{"./data/input2.json"}, []byte("[{\"id\":\"wqweq\",\"name\":\"test json 2\"},{\"name\":\"test json\",\"id\":\"jhasdad\"}]")},
		{"bar", args{"./data/input1.json"}, []byte("[{\"name\":\"test json\",\"id\":\"jhasdad\"},{\"name\":\"test json 2\",\"id\":\"wqweq\"}]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileJSONToByte(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileJSONToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
