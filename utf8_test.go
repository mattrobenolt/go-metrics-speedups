package speedups_test

import (
	"testing"

	"go.withmatt.com/metrics"
	_ "go.withmatt.com/metrics-speedups"
)

func TestImporting(t *testing.T) {
	metrics.MustIdent("utf8_test")
}
