package validation

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func MaxLengthChecker() Checker {
	return &maxLengthChecker{}
}

type maxLengthChecker struct{}

func (c *maxLengthChecker) Name() string {
	return "maxLength"
}

func (c *maxLengthChecker) Check(v interface{}, expects string) error {
	s, err := IsString(v)
	if err != nil {
		return err
	}

	max, err := IsStringInt(expects)
	if err != nil {
		return err
	}

	l := int64(utf8.RuneCountInString(s))
	if l > max {
		return errors.New(fmt.Sprintf(MaxLengthValueError, max))
	}

	return nil
}
