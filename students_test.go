package coverage

import (
	"os"
	// "errors"
	"reflect"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeople_Len(t *testing.T) {
	tests := []struct {
		name string
		p    People
		want int
	}{
		{name: "first", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"}}, want: 2},
		{name: "second", p: People{Person{firstName:"test_1",lastName:"test"}}, want: 1},
		{name: "third", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"},Person{firstName:"test_3",lastName:"test"}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("People.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPeople_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
		want bool
	}{
		{name: "first", p: People{Person{firstName:"test",lastName:"test", birthDay: time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},Person{firstName:"test",lastName:"test", birthDay: time.Date(1987, 4, 3, 0, 0, 0, 0, time.UTC)}}, args: args{i: 0, j: 1}, want: true},
		{name: "second", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"},Person{firstName:"test_3",lastName:"test"}}, args: args{i: 0, j: 2}, want: true},
		{name: "third", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"},Person{firstName:"test_3",lastName:"test"}}, args: args{i: 0, j: 1}, want: true},
		{name: "fourth", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"},Person{firstName:"test_3",lastName:"test"}}, args: args{i: 2, j: 0}, want: false},
		{name: "fifth", p: People{Person{firstName:"test",lastName:"test_1"},Person{firstName:"test",lastName:"test_2"},Person{firstName:"test",lastName:"test_3"}}, args: args{i: 2, j: 0}, want: false},
		{name: "sixth", p: People{Person{firstName:"test",lastName:"test_1"},Person{firstName:"test",lastName:"test_2"},Person{firstName:"test",lastName:"test_3"}}, args: args{i: 1, j: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("People.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeople_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
	}{
		{name: "first", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"}}, args: args{i: 0, j: 1}},
		{name: "second", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"},Person{firstName:"test_3",lastName:"test"}}, args: args{i: 0, j: 2}},
		{name: "third", p: People{Person{firstName:"test_1",lastName:"test"},Person{firstName:"test_2",lastName:"test"},Person{firstName:"test_3",lastName:"test"}}, args: args{i: 0, j: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}
func TestNew(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{name: "first", args: args{"1 2 3\n4 5 6 \n7 8 9"}, want: &Matrix{rows: 3, cols: 3, data: []int{1,2,3,4,5,6,7,8,9}}, wantErr: false},
		{name: "second", args: args{"1 2 \n3 4\n5 6"}, want: &Matrix{rows: 3, cols: 2, data: []int{1,2,3,4,5,6}}, wantErr: false},
		{name: "third", args: args{"1 2 7\n3 4\n5 6"}, want: nil, wantErr: true},
		{name: "fourth", args: args{"1 0 0 0\n0 1 0 0\n0 0 1 0\n0 0 0 1"}, want: &Matrix{rows: 4, cols: 4, data: []int{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}}, wantErr: false},
		{name: "fifth", args: args{"1 2 \n3 a\n5 6"}, want: nil, wantErr: true},
		{name: "fifth", args: args{"1 2 ф\n3 aы ы\n5 6 5"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want [][]int
	}{
		{name: "first", m: Matrix{rows: 3, cols: 3, data: []int{1,2,3,4,5,6,7,8,9}}, want: [][]int{{1,2,3},{4,5,6},{7,8,9}}},
		{name: "second", m: Matrix{rows: 3, cols: 2, data: []int{1,2,3,4,5,6}}, want: [][]int{{1,2,},{3,4},{5,6}}},
		{name: "third", m: Matrix{rows: 4, cols: 4, data: []int{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}}, want: [][]int{{1,0,0,0},{0,1,0,0},{0,0,1,0},{0,0,0,1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Rows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Cols(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want [][]int
	}{
		{name: "first", m: Matrix{rows: 3, cols: 3, data: []int{1,2,3,4,5,6,7,8,9}}, want: [][]int{{1,4,7},{2,5,8},{3,6,9}}},
		{name: "second", m: Matrix{rows: 3, cols: 2, data: []int{1,2,3,4,5,6}}, want: [][]int{{1,3,5},{2,4,6}}},
		{name: "third", m: Matrix{rows: 4, cols: 4, data: []int{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}}, want: [][]int{{1,0,0,0},{0,1,0,0},{0,0,1,0},{0,0,0,1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Cols(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	type args struct {
		row   int
		col   int
		value int
	}
	tests := []struct {
		name string
		m    *Matrix
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "first", m: &Matrix{rows: 3, cols: 3, data: []int{1,2,3,4,5,6,7,8,9}}, args: args{row: 0, col: 0, value: 1}, want: true},
		{name: "second", m: &Matrix{rows: 3, cols: 2, data: []int{1,2,3,4,5,6}}, args: args{row: 2, col: 1, value: 6}, want: true},
		{name: "third", m: &Matrix{rows: 4, cols: 4, data: []int{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}}, args: args{row: 2, col: 3, value: 0}, want: true},
		{name: "third", m: &Matrix{rows: 4, cols: 4, data: []int{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}}, args: args{row: -2, col: 3, value: 0}, want: false},
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Set(tt.args.row, tt.args.col, tt.args.value); got != tt.want {
				t.Errorf("Matrix.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}