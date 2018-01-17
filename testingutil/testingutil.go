package testingutil

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Assert fails the test if the condition is false
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.Fail()
	}
}

// Ok fails the test if an error is present
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n", filepath.Base(file), line, err.Error())
		tb.Fail()
	}
}

// Error fails the test if no errors is present
func Error(tb testing.TB, expected error, actual string) {
	if expected == nil || expected.Error() != actual {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: expected error %s\n\tgot: %s\033[39m\n", filepath.Base(file), line, expected.Error(), actual)
		tb.Fail()
	}
}

// Equals fails the test if the 'expected' value is not equal to 'actual' value
func Equals(tb testing.TB, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\texpected: %#v\n\tgot: %#v\033[39m\n", filepath.Base(file), line, expected, actual)
		tb.Fail()
	}
}
