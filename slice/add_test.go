package slice

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args[T any] struct {
		src   []T
		elm   T
		index int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ int]{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args[ /* TODO: Insert concrete types here */ int]{
				src:   []int{1, 2, 3, 4, 5},
				elm:   6,
				index: 2,
			},
			want:    []int{1, 2, 6, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "test2",
			args: args[ /* TODO: Insert concrete types here */ int]{
				src:   []int{1, 2, 3, 4, 5},
				elm:   6,
				index: 0,
			},
			want:    []int{6, 1, 2, 3, 4, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.src, tt.args.elm, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
