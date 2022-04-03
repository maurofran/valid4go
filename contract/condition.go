package contract

type defaultStringer string

func (def defaultStringer) String() string {
	return string(def)
}

const defaultMessageBuilder = defaultStringer("")

// Condition is the interface used to check for contract conditions.
type Condition interface {
	IsNotSatisfied() bool
	Message() string
}
