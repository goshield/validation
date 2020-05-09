package validation

import (
	"errors"
	"reflect"
)

// NilChecker returns a checker
func NilChecker() Checker {
	return &nilChecker{}
}

type nilChecker struct{}

func (c *nilChecker) Name() string {
	return "nil"
}

func (c *nilChecker) Check(v interface{}, expects string) error {
	if expects == False && reflect.ValueOf(v).IsNil() {
		return errors.New(NilValueError)
	}

	return nil
}
