package main

import "testing"

func Test_permutationsForString(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trivial",
			args: args{s: "of"},
			want: 2,
		},
		{
			name: "trivial",
			args: args{s: "for"},
			want: 6,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := permutationsForString(tt.args.s); got != tt.want {
				t.Errorf("permutationsForString() = %v, want %v", got, tt.want)
			}
		})
	}
}
