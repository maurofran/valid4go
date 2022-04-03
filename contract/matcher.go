package contract

import (
	"github.com/maurofran/hamcrest4go/matcher"
	"github.com/maurofran/valid4go/internal/message"
)

type matcherCondition[T any] struct {
	value   T
	matcher matcher.Matcher[T]
}

func (m matcherCondition[T]) IsNotSatisfied() bool {
	return !m.matcher.Matches(m.value)
}

func (m matcherCondition[T]) Message() string {
	return message.DescribingMismatchOf[T](m.value, m.matcher).String()
}

func Matching[T any](value T, m matcher.Matcher[T]) Condition {
	return matcherCondition[T]{
		value:   value,
		matcher: m,
	}
}
