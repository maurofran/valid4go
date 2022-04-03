package argument

import (
	"github.com/maurofran/hamcrest4go/described"
	"github.com/maurofran/hamcrest4go/matcher"
)

// Named is a matcher used to add an argument name to a matcher description.
func Named[T any](name string, matcher matcher.Matcher[T]) matcher.Matcher[T] {
	return described.As[T]("%0 = %1", matcher, name, matcher)
}
