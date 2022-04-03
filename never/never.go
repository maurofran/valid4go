package never

import "github.com/maurofran/valid4go/internal/message"

var (
	noError   error  = nil
	noMessage string = ""
)

// GetHere validates that unreachable code have been reached.
// This is considered to be a programming error.
func GetHere() {
	GetHereErrorf(noError, noMessage)
}

// GetHeref validates that unreachable code have been reached.
// This is considered to be a programming error.
func GetHeref(msg string, args ...interface{}) {
	GetHereErrorf(noError, msg, args...)
}

// GetHereError validates that unreachable code have been reached.
// This is considered to be a programming error.
func GetHereError(err error) {
	GetHereErrorf(err, noMessage)
}

// GetHereErrorf validates that unreachable code have been reached.
// This is considered to be a programming error.
func GetHereErrorf(err error, msg string, args ...interface{}) {
	policy(err, message.Describing(msg, args...))
}
