// Package goerror stores generated error
package goerror

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// GoError structure supports the goerror structure
type GoError struct {
	err    error  // system error generated
	source string // filename
	lineno int    // lineno in file causing error
	status string // error status  "FATAL", "ERROR", "WARNING"
	arg    int    // error number
	msg    string // error message
}

// lastError is a global error
var lastError = GoError{}

// GetLastError function returns the last goerror as a formatted string
func GetLastError() string {
	return lastError.Error()
}

// LastError returns the last global goerror set
func LastError() GoError {
	return lastError
}

// PrintLastError function prints the code goerror message to stdout
func PrintLastError() {
	fmt.Printf(lastError.Error())
}

// SetLastError function sets the global goerror structure with an error
func SetLastError(err error) {
	lastError.err = err
	lastError.source = ""
	lastError.lineno = 0
	lastError.status = "ERROR"
	lastError.arg = 0
	lastError.msg = fmt.Sprint(err)
}

// FromError function creates a new goerror structure and initialises it with a go error
func FromError(err error) *GoError {
	msg := fmt.Sprint(err)
	return &GoError{err: err, source: "", lineno: 0, status: "ERROR", arg: 0, msg: msg}
}

// New function creates a new goerror structure and initialises it from a fromatted string
// example source error string "main.go:205:ERROR:0200:Could not convert string to int"
// example error string "ERROR:0012:Incorrect username or password!"
func New(errstr string) *GoError {
	err := errors.New("")
	s := strings.Split(errstr, ":")
	if len(s) == 3 {
		status := s[0]
		arg := toInt(s[1])
		msg := s[2]
		return &GoError{err: err, source: "", lineno: 0, status: status, arg: arg, msg: msg}
	} else if len(s) == 5 {
		source := s[0]
		lineno := toInt(s[1])
		status := s[2]
		arg := toInt(s[3])
		msg := s[4]
		return &GoError{err: err, source: source, lineno: lineno, status: status, arg: arg, msg: msg}
	}
	return &GoError{}
}

// SetCodeError function updates the goerror object
func (e *GoError) SetCodeError(source string, lineno int, status string, arg int, msg string) {
	if e.err == nil {
		e.err = errors.New("")
	}
	e.source = source
	e.lineno = lineno
	e.status = status
	e.arg = arg
	e.msg = msg
}

// CodeError function returns the goerror in code format as a string
func (e *GoError) CodeError() string {
	return fmt.Sprintf("%s:file %s:line %d: error %d: %s", e.status, e.source, e.lineno, e.arg, e.msg)
}

// Error function prints the go error source if it exists
func (e *GoError) Error() string {
	if fmt.Sprint(e.err) == "" {
		return "no external package error"
	}
	return fmt.Sprint(e.err)
}

// IsNil function returns false if there is err, else returns true
func (e *GoError) IsNil() bool {
	if e.err != nil {
		return false
	}
	return true
}

// PrintError function prints the user goerror message to stdout
func (e *GoError) PrintError() {
	fmt.Printf("%s:%d:%s\n", e.status, e.arg, e.msg)
}

// PrintCodeError function prints the code goerror message to stdout
func (e *GoError) PrintCodeError() {
	fmt.Printf("%s %4d: %s:%05d...%s\n", e.status, e.arg, e.source, e.lineno, e.msg)
}

// UserError function returns the user goerror message
func (e *GoError) UserError() string {
	return fmt.Sprintf("%s:%d:%s", e.status, e.arg, e.msg)
}

// CODE function returns the code goerror message
func (e *GoError) CODE() string {
	return fmt.Sprintf("%s %4d: %s:%05d...%s", e.status, e.arg, e.source, e.lineno, e.msg)
}

// USER function returns the user goerror message
func (e *GoError) USER() string {
	return fmt.Sprintf("%s:%d:%s", e.status, e.arg, e.msg)
}

// SetSource function sets the source file name in the error
func (e *GoError) SetSource(source string) {
	e.source = source
}

// SetLineNo function sets the line number of the source file in the error
func (e *GoError) SetLineNo(lineno int) {
	e.lineno = lineno
}

// SetStatus function sets the status of the error
func (e *GoError) SetStatus(status string) {
	e.status = status
}

// SetArg function sets the error number in the error
func (e *GoError) SetArg(arg int) {
	e.arg = arg
}

// SetMsg function sets the message in the error
func (e *GoError) SetMsg(msg string) {
	e.msg = msg
}

// Status function returns the status of the error
func (e *GoError) Status() string {
	return e.status
}

// Msg function returns the message in the error
func (e *GoError) Msg() string {
	return e.msg
}

// Arg function returns the error number in the error
func (e *GoError) Arg() int {
	return e.arg
}

// Source function returns the source file name in the error
func (e *GoError) Source() string {
	return e.source
}

// LineNo function returns the line number of the source file in the error
func (e *GoError) LineNo() int {
	return e.lineno
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		return i
	}
	return 0
}
