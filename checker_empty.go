package validation

import "errors"

func EmptyChecker() Checker {
	return &emptyChecker{}
}

type emptyChecker struct {}

func (c *emptyChecker) Name() string {
	return "empty"
}

func (c *emptyChecker) Check(v interface{}, expects string) error {
	if v == "" && expects == False {
		return errors.New(EmptyValueError)
	}

	return nil
}


