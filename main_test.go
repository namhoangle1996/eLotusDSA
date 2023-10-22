package main

import (
	"reflect"
	"testing"
)

func Test_grayCode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "",
			args: args{
				n: 1,
			},
			wantRes: []int{0, 1},
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			wantRes: []int{0, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := grayCode(tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("grayCode() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "",
			args: args{
				nums1: []int{0, 1, 2, 3, 4},
				nums2: []int{0, 2, 3, 5},
			},
			wantRes: 3,
		},
		{
			name: "",
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0, 0, 0},
			},
			wantRes: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findLength(tt.args.nums1, tt.args.nums2); gotRes != tt.wantRes {
				t.Errorf("findLength() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
