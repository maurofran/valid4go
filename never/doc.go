// Package never is the entry point for unreachable code validations.
//
// Violations indicate the presence of a programming error.
// Never try to recover from a programming error. Rather a reasonable global approach would be to
// return a HTTP 500 response, restart the application or similar.
//
// Make use of hamcrest matchers to conveniently express conditions and get clear error messages.
// Example:
//
//		require.That[int32](i, number.GreaterThan(0))
//
// All error messages are formatted string as defined in fmt.Sprintf function.
// Example:
//
// 		require.Conditionf(i > 0, "The value must be greater than zero: %d", i)
package never
