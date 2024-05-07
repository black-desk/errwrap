package errwrap_test

import (
	"fmt"

	"github.com/black-desk/errwrap"
)

func ExampleWrap_with_annotation() {
	fn := func(x int) (err error) {
		defer errwrap.Wrap(&err, "fn(%d)", x)
		return fmt.Errorf("error")
	}
	err := fn(1)
	fmt.Println(err)

	// Output: fn(1): error
}

func ExampleWrap_without_annotation() {

	fn := func() (err error) {
		defer errwrap.Wrap(&err)
		return fmt.Errorf("error")
	}
	err := fn()
	fmt.Println(err)

	// Output: error
}
