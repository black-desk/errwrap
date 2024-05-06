package errwrap_test

import (
	"fmt"

	"github.com/black-desk/errwrap"
)

func ExampleWrap() {
	var err error
	err = fmt.Errorf("some error")
	errwrap.Wrap(&err, "wrapped")
	fmt.Println(err)
	// Output: wrapped: some error
}
