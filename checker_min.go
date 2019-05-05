package validation

import (
	"errors"
	"fmt"
	"strconv"
)

func MinChecker() Checker {
	return &minChecker{}
}

type minChecker struct{}

func (c *minChecker) Name() string {
	return "min"
}

// Check tests input value with expectation
func (c *minChecker) Check(v interface{}, expects string) error {
	if err := IsNumber(v); err != nil {
		return err
	}

	ei, err := strconv.ParseInt(expects, 10, 64)
	if err == nil {
		return c.checkInt(v, ei)
	}

	ef, err := strconv.ParseFloat(expects, 64)
	if err == nil {
		return c.checkFloat(v, ef)
	}

	return errors.New(NotNumberValueError)
}

func (c *minChecker) checkInt(v interface{}, expects int64) error {
	i, err := IsInt(v)
	if err != nil {
		return err
	}

	if i < expects {
		return errors.New(fmt.Sprintf(MinValueError, expects))
	}

	return nil
}

func (c *minChecker) checkFloat(v interface{}, expects float64) error {
	f, err := IsFloat(v)
	if err != nil {
		return err
	}

	if f < expects {
		return errors.New(fmt.Sprintf(MinValueError, expects))
	}

	return nil
}
