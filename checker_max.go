package validation

import (
	"errors"
	"fmt"
	"strconv"
)

func MaxChecker() Checker {
	return &maxChecker{}
}

type maxChecker struct{}

func (c *maxChecker) Name() string {
	return "max"
}

// Check tests input value with expectation
func (c *maxChecker) Check(v interface{}, expects string) error {
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

func (c *maxChecker) checkInt(v interface{}, expects int64) error {
	i, err := IsInt(v)
	if err != nil {
		return err
	}

	if i > expects {
		return errors.New(fmt.Sprintf(MaxValueError, expects))
	}

	return nil
}

func (c *maxChecker) checkFloat(v interface{}, expects float64) error {
	f, err := IsFloat(v)
	if err != nil {
		return err
	}

	if f > expects {
		return errors.New(fmt.Sprintf(MaxValueError, expects))
	}

	return nil
}
