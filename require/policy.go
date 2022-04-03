package require

import "github.com/maurofran/valid4go/contract"

// Policy is the policy used to check the preconditions.
type Policy func(condition contract.Condition)

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

func defaultPolicy(condition contract.Condition) {
	if condition.IsNotSatisfied() {
		panic(condition.Message())
	}
}

func disabledPolicy(condition contract.Condition) {
	// Do nothing.
}
