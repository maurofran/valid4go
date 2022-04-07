package validate

import (
	"fmt"
	"github.com/maurofran/valid4go/contract"
)

// Policy is the policy used to check the preconditions.
type Policy func(condition contract.Condition, err error) error

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

func defaultPolicy(condition contract.Condition, err error) error {
	if condition.IsNotSatisfied() {
		return fmt.Errorf("%w: %s", err, condition.Message())
	}
	return nil
}

func disabledPolicy(condition contract.Condition, err error) error {
	return nil
}
