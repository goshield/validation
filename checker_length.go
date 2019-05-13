package validation

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func LengthChecker() Checker {
	return &lengthChecker{}
}

type lengthChecker struct{}

func (c *lengthChecker) Name() string {
	return "length"
}

func (c *lengthChecker) Check(v interface{}, expects string) error {
	s, err := IsString(v)
	if err != nil {
		return err
	}

	length, err := IsStringInt(expects)
	if err != nil {
		return err
	}

	l := int64(utf8.RuneCountInString(s))
	if l != length {
		return errors.New(fmt.Sprintf(LengthValueError, length))
	}

	return nil
}
