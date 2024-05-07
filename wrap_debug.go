//go:build debug
// +build debug

package errwrap

import (
	"fmt"
	"runtime"
)

func formatAnnotate(annotate []any) string {
	if len(annotate) == 0 {
		return ""
	}

	msg, ok := annotate[0].(string)
	if !ok {
		panic("annotate should start with string")
	}

	return fmt.Sprintf("\t"+msg+"\n", annotate[1:]...)
}

func Wrap(err *error, annotate ...any) {
	if err == nil || *err == nil {
		return
	}

	msg := formatAnnotate(annotate)

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("failed to get caller")
	}

	*err = fmt.Errorf(
		"%s:%d (%x)%s\n%w",
		file, line, pc, msg,
		*err,
	)
}
