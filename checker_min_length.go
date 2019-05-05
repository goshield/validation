package validation

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func MinLengthChecker() Checker {
	return &minLengthChecker{}
}

type minLengthChecker struct{}

func (c *minLengthChecker) Name() string {
	return "minLength"
}

func (c *minLengthChecker) Check(v interface{}, expects string) error {
	s, err := IsString(v)
	if err != nil {
		return err
	}

	min, err := IsStringInt(expects)
	if err != nil {
		return err
	}

	l := int64(utf8.RuneCountInString(s))
	if l < min {
		return errors.New(fmt.Sprintf(MinLengthValueError, min))
	}

	return nil
}
