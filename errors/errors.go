// Package errors implements functions to manipulate errors.
package errors

// Error is a trivial implementation of error.
type Error string

func (e Error) Error() string {
	return string(e)
}

// New returns an error that formats as the given text.
func New(text string) Error {
	return Error(text)
}
