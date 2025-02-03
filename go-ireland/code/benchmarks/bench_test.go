package bench_test

import (
	"testing"
)

// STARTEST OMIT
func TestMyFunc(t *testing.T) { // HL11
	// test logic
}

//ENDTEST OMIT

// START OMIT

func BenchmarkMyFunc(b *testing.B) { // HL12
	for i := 0; i < b.N; i++ { // HL13
		MyFunc()
	}
}

// END OMIT

func MyFunc() {}

func Bar(s string) string {
	return s
}

// STARTTABLE OMIT
var tt = []string{"uno", "due", "tre"}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tt {
			Bar(tc)
		}
	}
}

// ENDTABLE OMIT
