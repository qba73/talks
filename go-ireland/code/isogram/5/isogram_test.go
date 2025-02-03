package isogram_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/isogram"
)

func TestIsIsogram(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := isogram.IsIsogram(tc.input)
			if !cmp.Equal(tc.want, got) {
				t.Error(cmp.Diff(tc.want, got))
			}
		})
	}
}

var tt = []struct {
	desc  string
	input string
	want  bool
}{
	{
		desc:  "isogram with only lower case characters",
		input: "isogram",
		want:  true,
	},
}

// STARTISO OMIT
var Res bool // HL

func BenchmarkIsIsogram(b *testing.B) {
	var got bool // HL
	for i := 0; i < b.N; i++ {
		for _, tc := range tt {
			got = isogram.IsIsogram(tc.input) // HL
		}
	}
	Res = got // HL
}

// ENDISO OMIT
