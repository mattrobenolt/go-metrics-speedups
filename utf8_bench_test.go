package speedups_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"go.withmatt.com/metrics-speedups"
)

func BenchmarkUTF8Valid(b *testing.B) {
	long := strings.Repeat("foo", 100)
	short := "foobar"
	for i, tc := range []string{long, short} {
		suffix := "long"
		if i == 1 {
			suffix = "short"
		}
		b.Run("stdlib-"+suffix, func(b *testing.B) {
			for b.Loop() {
				utf8.ValidString(tc)
			}
		})

		b.Run("asm-"+suffix, func(b *testing.B) {
			for b.Loop() {
				speedups.UTF8ValidString(tc)
			}
		})
	}
}
