package arraypractise

import (
	"reflect"
	"testing"
)

// func TestIntMin(t *testing.T) {
// 	min := IntMin(1, 5)
// 	if min != 1 {
// 		t.Errorf("Expected = %d; got = %d", 1, min)
// 	}
// }

func TestIntMinWithTable(t *testing.T) {
	type args struct {
		a, b int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				a: 1,
				b: 5,
			},
			want: 1,
		},
		{
			name: "test 1",
			args: args{
				a: 10,
				b: 11,
			},
			want: 10,
		},
		{
			name: "test 1",
			args: args{
				a: 1,
				b: -5,
			},
			want: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min := IntMin(tt.args.a, tt.args.b)
			// if min != tt.want {
			// 	t.Errorf("Expected = %d; got = %d", tt.want, min)
			// }
			if !reflect.DeepEqual(min, tt.want) {
				t.Errorf("Expected = %d; got = %d", tt.want, min)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 3)
	}
}

func TestCyclicRotate(t *testing.T) {
	type args struct {
		a []int
		k int
	}
	var tests = []struct {
		name      string
		args      args
		want      []int
		wantError string
	}{
		{
			name: "test1",
			args: args{
				a: []int{1, 2, 3},
				k: -1,
			},
			wantError: " K cannot be negative",
		},
		{
			name: "test2",
			args: args{
				a: []int{1, 2, 3},
				k: 8,
			},
			want: []int{2, 3, 1},
		},
		{
			name: "Test a",
			args: args{
				a: []int{3, 8, 9, 7, 6},
				k: 1,
			},
			want: []int{6, 3, 8, 9, 7},
		},
		{
			name: "Test b",
			args: args{
				a: []int{3, 8, 9, 7, 6},
				k: 3,
			},
			want: []int{9, 7, 6, 3, 8},
		},
		{
			name: "Test c",
			args: args{
				a: []int{3, 8, 9, 7, 6},
				k: 0,
			},
			want: []int{3, 8, 9, 7, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testing, err := CyclicRotate(tt.args.a, tt.args.k)
			if err != nil && !reflect.DeepEqual(err.Error(), tt.wantError) {
				t.Errorf("Expected = %v: got = %v\n", tt.wantError, err.Error())
			} else if !reflect.DeepEqual(testing, tt.want) {
				t.Errorf("Expected = %d: got = %d\n", tt.want, testing)
			}
		})
	}
}

func BenchmarkCyclicRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CyclicRotate([]int{1, 2, 3, 4, 5, 6, 7}, 10)
	}
}
