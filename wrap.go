//go:build !debug
// +build !debug

package errwrap

import (
	"fmt"
)

func formatAnnotate(annotate []any) string {
	if len(annotate) == 0 {
		return ""
	}

	msg, ok := annotate[0].(string)
	if !ok {
		panic("annotate should start with string")
	}

	return fmt.Sprintf(msg, annotate[1:]...)
}

// You can use this Wrap function to wrap an error with an annotation.
// The annotation will be passed to fmt.Sprintf,
// so it must starts with a format string.
//
// Add build tag `debug` to
// make wrapped error message cantains file name, line number and pc.
func Wrap(err *error, annotate ...any) {
	if err == nil || *err == nil {
		return
	}

	msg := formatAnnotate(annotate)

	if msg != "" {
		*err = fmt.Errorf(msg+": %w", *err)
	}

	return
}
