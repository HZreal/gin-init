package main

import "testing"

/**
 * @Author elastic·H
 * @Date 2024-08-09
 * @File: client_test.go
 * @Description:
 */

func Test_invoke(t *testing.T) {
	type args struct {
		x int32
		y int32
	}
	tests := []struct {
		name         string
		args         args
		wantSum      int32
		wantSubtract int32
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{x: 10, y: 20}, wantSum: 30, wantSubtract: -10},
		{name: "case2", args: args{x: 10, y: 5}, wantSum: 15, wantSubtract: 5},
		{name: "case3", args: args{x: 20, y: 20}, wantSum: 40, wantSubtract: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotSubtract := invoke(tt.args.x, tt.args.y)
			if gotSum != tt.wantSum {
				t.Errorf("invoke() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotSubtract != tt.wantSubtract {
				t.Errorf("invoke() gotSubtract = %v, want %v", gotSubtract, tt.wantSubtract)
			}
		})
	}
}
