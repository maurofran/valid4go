// Package validate is the entry point for global preconditions.
//
// Use the function of this package to return recoverable errors if validation fails.
//
// Make use of hamcrest matchers to conveniently express conditions and get clear error messages.
// Example:
//
//		validate.That[int32](i, number.GreaterThan(0), ErrInvalidArgument)
//
// All error messages are formatted string as defined in fmt.Sprintf function.
// Example:
//
// 		validate.Conditionf(i > 0, ErrInvalidArgument, "The value must be greater than zero: %d", i)
package validate
