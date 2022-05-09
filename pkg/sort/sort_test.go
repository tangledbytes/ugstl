package sort

import (
	"math/rand"
	stdSort "sort"
	"testing"

	"github.com/utkarsh-pro/ugstl/pkg/types"
)

func TestSelection(t *testing.T) {
	type args struct {
		arr  []int
		less types.Less[int]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sort random array of size 1000",
			args: args{
				arr: rand.Perm(1000),
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
		{
			name: "Sort already sorted array",
			args: args{
				arr: []int{1, 2, 3, 4},
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
		{
			name: "Sort empty array",
			args: args{
				arr: []int{},
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
		{
			name: "Sort reverse sorted array",
			args: args{
				arr: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Selection(tt.args.arr, tt.args.less)

			if !stdSort.IntsAreSorted(tt.args.arr) {
				t.Error("expected array to be sorted: ", tt.args.arr)
			}
		})
	}
}

func TestHeap(t *testing.T) {
	type args struct {
		arr  []int
		less types.Less[int]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sort random array of size 1000",
			args: args{
				arr: rand.Perm(1000),
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
		{
			name: "Sort already sorted array",
			args: args{
				arr: []int{1, 2, 3, 4},
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
		{
			name: "Sort empty array",
			args: args{
				arr: []int{},
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
		{
			name: "Sort reverse sorted array",
			args: args{
				arr: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
				less: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Heap(tt.args.arr, tt.args.less)

			if !stdSort.IntsAreSorted(tt.args.arr) {
				t.Error("expected array to be sorted: ", tt.args.arr)
			}
		})
	}
}
