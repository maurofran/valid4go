package contract

import "fmt"

type booleanCondition struct {
	isSatisfied    bool
	messageBuilder fmt.Stringer
}

func (b booleanCondition) IsNotSatisfied() bool {
	return !b.isSatisfied
}

func (b booleanCondition) Message() string {
	return b.messageBuilder.String()
}

func Boolean(isSatisfied bool, messageBuilder fmt.Stringer) Condition {
	if messageBuilder == nil {
		messageBuilder = defaultMessageBuilder
	}
	return booleanCondition{
		isSatisfied:    isSatisfied,
		messageBuilder: messageBuilder,
	}
}
