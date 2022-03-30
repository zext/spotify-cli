package stringsx

import "testing"

func TestJoinAnd(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0 element",
			args: args{
				s: []string{},
			},
			want: "",
		},
		{
			name: "1 element",
			args: args{
				s: []string{"a"},
			},
			want: "a",
		},
		{
			name: "2 element",
			args: args{
				s: []string{"a", "b"},
			},
			want: "a and b",
		},
		{
			name: "3 element",
			args: args{
				s: []string{"a", "b", "c"},
			},
			want: "a, b, and c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinAnd(tt.args.s); got != tt.want {
				t.Errorf("JoinAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
