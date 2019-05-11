package validation

import (
	"errors"
	"fmt"
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
	sv := fmt.Sprintf("%v", v)

	return c.checkString(sv, in)
}

func (c *inChecker) checkString(v string, in []string) error {
	for _, s := range in {
		if strings.Compare(v, s) == 0 {
			return nil
		}
	}

	return errors.New(fmt.Sprintf(ItemNotFoundInListError, v, in))
}
