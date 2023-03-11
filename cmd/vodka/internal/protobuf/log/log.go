//go:build debug
// +build debug

package log

import (
	"fmt"
	"os"
)

func Debug(format string, fields ...interface{}) {
	fmt.Fprintf(os.Stderr, format, fields...)
}
