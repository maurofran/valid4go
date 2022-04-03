package message

import (
	"fmt"
	"github.com/maurofran/hamcrest4go/matcher"
	"strings"
)

type Message string

func (m Message) String() string {
	return string(m)
}

func Describing(msg string, args ...interface{}) Message {
	return Message(withFormattedMessage(msg, args...))
}

func withFormattedMessage(msg string, args ...interface{}) string {
	if msg != "" {
		return fmt.Sprintf(msg, args...)
	}
	if len(args) > 0 {
		return fallbackFormattingOf("null", args...)
	}
	return ""
}

func fallbackFormattingOf(msg string, args ...interface{}) string {
	var builder strings.Builder
	builder.WriteString(msg)
	for i, arg := range args {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(fmt.Sprintf("%s", arg))
	}
	return builder.String()
}

func DescribingMismatchOf[T any](value T, m matcher.Matcher[T]) Message {
	return Message(withMismatchMessageOf[T](value, m))
}

func withMismatchMessageOf[T any](value T, m matcher.Matcher[T]) string {
	description := matcher.StringDescription()
	description.AppendText("expected: ")
	description.AppendDescriptionOf(m)
	description.AppendText("\n but: ")
	m.DescribeTo(description)
	return description.String()
}
