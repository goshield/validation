package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func InChecker() Checker {
	return new(inChecker)
}

type inChecker struct{}

func (c *inChecker) Name() string {
	return "in"
}

func (c *inChecker) Check(v interface{}, expects string) error {
	if expects == "" {
		return errors.New(fmt.Sprintf(EmptyListFoundError, v))
	}
	in := strings.Split(expects, ",")

	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int64:
		return c.checkInt(reflect.ValueOf(v).Int(), in)
	case reflect.Float64:
		return c.checkFloat(reflect.ValueOf(v).Float(), in)
	case reflect.String:
		return c.checkString(reflect.ValueOf(v).String(), in)
	default:
		return errors.New(InvalidArgumentError)
	}
}

func (c *inChecker) checkInt(v int64, in []string) error {
	for _, s := range in {
		i, err := IsStringInt(s)
		if err != nil {
			return err
		}
		if v == i {
			return nil
		}
	}

	return errors.New(fmt.Sprintf(ItemNotFoundInListError, v, in))
}

func (c *inChecker) checkFloat(v float64, in []string) error {
	for _, s := range in {
		i, err := IsStringFloat(s)
		if err != nil {
			return err
		}
		if v == i {
			return nil
		}
	}

	return errors.New(fmt.Sprintf(ItemNotFoundInListError, v, in))
}

func (c *inChecker) checkString(v string, in []string) error {
	for _, s := range in {
		if strings.Compare(v, s) == 0 {
			return nil
		}
	}

	return errors.New(fmt.Sprintf(ItemNotFoundInListError, v, in))
}
