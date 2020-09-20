package err

import (
	"fmt"
	// "strconv"
	"strings"
)

//Error const
const (
	PermissionDenied  = "permission denied"
	InvalidParameters = "invalid perameters"
)

const (
	CodePermissionDenied = 1000*1 + iota
	CodeInvalidParameters
)

//Custom Error
var (
	ErrPermissionDenied  = wrapError{code: CodePermissionDenied, msg: PermissionDenied}
	ErrInvalidParameters = wrapError{code: CodeInvalidParameters, msg: InvalidParameters}
)

type wrapError struct {
	err  error
	code int
	msg  string
}

func (err wrapError) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%v %s: %v", err.code, err.msg, err.err)
	}
	return err.msg
}

func (err wrapError) wrap(inner error) error {
	return wrapError{code: err.code, msg: err.msg, err: inner}
}

func (err wrapError) Unwrap() error {
	return err.err
}

func (err wrapError) Is(target error) bool {
	ts := target.Error()
	return ts == err.msg || strings.HasPrefix(ts, err.msg+": ")
}
