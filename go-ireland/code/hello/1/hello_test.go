package hello_test

import (
	"testing"

	"github.com/qba73/hello"
)

var tt = []struct {
	desc  string
	input string
	want  string
}{
	{
		desc:  "empty name",
		input: "",
		want:  "One for you, one for me.",
	},
	{
		desc:  "name is Alice",
		input: "Alice",
		want:  "One for Alice, one for me.",
	},
	{
		desc:  "name is Bob",
		input: "Bob",
		want:  "One for Bob, one for me.",
	},
}

func TestShareWith(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := hello.ShareWith(tc.input)
			if got != tc.want {
				t.Errorf("ShareWith(%q) = %q, want: %q", tc.input, got, tc.want)
			}
		})
	}
}

var result string

func BenchmarkShareWith(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		for _, tc := range tt {
			r = hello.ShareWith(tc.input)
		}
	}
	result = r
}
