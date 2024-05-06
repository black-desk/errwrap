package errwrap_test

import (
	"fmt"

	"github.com/black-desk/errwrap"
)

func ExampleWrap() {
	fn := func(x int) (err error) {
		defer errwrap.Wrap(&err, "fn(%d)", x)
		return fmt.Errorf("error")
	}
	err := fn(1)
	fmt.Println(err)
	// Output: fn(1): error
}
