package never

import "fmt"

var policy = defaultPolicy

// SetPolicy set the policy to be used in never package.
func SetPolicy(newPolicy Policy) {
	if policy == nil {
		policy = defaultPolicy
	} else {
		policy = newPolicy
	}
}

// Enable the never package assertions.
func Enable() {
	policy = defaultPolicy
}

// Disable the never package assertions.
func Disable() {
	policy = disabledPolicy
}

// Policy is the policy for unreachable code.
type Policy func(err error, messageBuilder fmt.Stringer)

func defaultPolicy(err error, messageBuilder fmt.Stringer) {
	panic(fmt.Errorf("%w: %s", err, messageBuilder.String()))
}

func disabledPolicy(err error, messageBuilder fmt.Stringer) {
	// Do nothing
}
