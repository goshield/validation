package validation

import (
	"errors"
	"fmt"
	"regexp"
)

func RegExpChecker() Checker {
	return &regExpChecker{}
}

type regExpChecker struct{}

func (c *regExpChecker) Name() string {
	return "regexp"
}

func (c *regExpChecker) Check(v interface{}, expects string) error {
	r, e := regexp.Compile(expects)
	if e != nil {
		return errors.New(fmt.Sprintf(InvalidRegexPatternError, expects))
	}

	s, err := IsString(v)
	if err != nil {
		return err
	}

	if r.MatchString(s) == false {
		return errors.New(fmt.Sprintf(RegexValueError, expects))
	}

	return nil
}
