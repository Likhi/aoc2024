package aoc2024

import "testing"

func Test_identical(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			args: args{
				a: []string{"X", "M", "A", "S"},
				b: []string{"X", "M", "A", "S"},
			},
			want: true,
		},
		{
			name: "mismatch",
			args: args{
				a: []string{"X", "M", "A", "S"},
				b: []string{"X", "M", "A", "F"},
			},
			want: false,
		},
		{
			name: "different_length",
			args: args{
				a: []string{"X", "M", "A", "S"},
				b: []string{"X", "M", "A"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := identical(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("identical() = %v, want %v", got, tt.want)
			}
		})
	}
}
