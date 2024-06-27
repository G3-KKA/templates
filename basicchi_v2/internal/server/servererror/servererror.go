package servererror

import "strings"

type ServerError struct {
	err  error
	hint string
}

func New(err error, hints ...string) ServerError {
	return ServerError{err: err, hint: strings.Join(hints, " ")}
}
func (se ServerError) Error() string {
	if se.hint == "" {
		return se.err.Error()
	}
	return se.err.Error() + " . Hints: " + se.hint
}
