package ensure

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/maurofran/valid4go/contract"
	"github.com/maurofran/valid4go/internal/message"
)

// Condition validates a precondition that clients are required to fulfill.
// Violations are considered to be programming errors on the clients part.
func Condition(condition bool) {
	policy(contract.Boolean(condition, nil))
}

// Conditionf validates a precondition that clients are required to fulfill.
// Violations are considered to be programming errors on the clients part.
func Conditionf(condition bool, msg string, args ...interface{}) {
	policy(contract.Boolean(condition, message.Describing(msg, args...)))
}

// That validates a precondition that clients are required to fulfill.
// Violations are considered to be programming errors on the clients part.
func That[T any](value T, matcher matcher.Matcher[T]) {
	policy(contract.Matching[T](value, matcher))
}
