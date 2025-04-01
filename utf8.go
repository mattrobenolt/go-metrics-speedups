package speedups

import (
	"unsafe"

	"github.com/segmentio/asm/utf8"
	_ "go.withmatt.com/metrics"
)

func init() {
	speedupsHook(UTF8ValidString)
}

//go:linkname speedupsHook go.withmatt.com/metrics.replaceUTF8ValidStringHook
func speedupsHook(func(string) bool)

func UTF8ValidString(s string) bool {
	b := unsafe.Slice(unsafe.StringData(s), len(s))
	return utf8.Validate(b).IsUTF8()
}
