package validate

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/maurofran/valid4go/contract"
	"github.com/maurofran/valid4go/internal/message"
)

// Condition validates a precondition that clients are required to fulfill.
// Violations are considered to be programming errors on the clients part.
func Condition(condition bool, err error) error {
	return policy(contract.Boolean(condition, nil), err)
}

// Conditionf validates a precondition that clients are required to fulfill.
// Violations are considered to be programming errors on the clients part.
func Conditionf(condition bool, err error, msg string, args ...interface{}) error {
	return policy(contract.Boolean(condition, message.Describing(msg, args...)), err)
}

// That validates a precondition that clients are required to fulfill.
// Violations are considered to be programming errors on the clients part.
func That[T any](value T, err error, matcher matcher.Matcher[T]) error {
	return policy(contract.Matching[T](value, matcher), err)
}
