package validation

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Validator interface {
	// Validate checks input value for error
	Validate(v interface{}) error

	// Extend registers a checker
	Extend(checker Checker) Validator
}

func New() Validator {
	return NewWithTag("validate")
}

// NewWithTag allows to create validator with custom tag rather than "validate"
func NewWithTag(tag string) Validator {
	v := &factoryValidator{
		tag:      tag,
		errorTag: "error",
		checkers: make(map[string]Checker),
	}

	return initializeCheckers(v)
}

func initializeCheckers(v Validator) Validator {
	return v.Extend(MinChecker()).
		Extend(MaxChecker()).
		Extend(MinLengthChecker()).
		Extend(MaxLengthChecker()).
		Extend(RangeChecker()).
		Extend(EmailChecker()).
		Extend(RegExpChecker()).
		Extend(InChecker()).
		Extend(EmptyChecker())
}

type factoryValidator struct {
	tag      string
	errorTag string
	checkers map[string]Checker
}

func (v *factoryValidator) Extend(checker Checker) Validator {
	v.checkers[checker.Name()] = checker
	return v
}

func (v *factoryValidator) Validate(input interface{}) error {
	t, err := v.validateType(input)
	if err != nil {
		return err
	}

	n := t.NumField()
	if n == 0 {
		// No fields => no process
		return nil
	}

	val, _ := v.valueOf(input)
	for i := 0; i < n; i++ {
		vf := val.Field(i)
		if vf.CanInterface() && (vf.Kind() == reflect.Struct || vf.Kind() == reflect.Ptr) {
			if err := v.Validate(vf.Interface()); err != nil {
				return err
			}
		}

		sf := t.Field(i)
		if _, ok := sf.Tag.Lookup(v.tag); ok == false {
			continue
		}

		tag := sf.Tag.Get(v.tag)
		if IsEmpty(tag) {
			continue
		}

		for k, p := range v.parseTags(tag) {
			c, ok := v.checkers[k]
			if ok == false {
				continue
			}
			if err := c.Check(vf.Interface(), p); err != nil {
				errMsg := sf.Tag.Get(v.errorTag)
				if !IsEmpty(errMsg) {
					return errors.New(errMsg)
				}
				return err
			}
		}
	}

	return nil
}

func (v *factoryValidator) validateType(input interface{}) (reflect.Type, error) {
	t := reflect.TypeOf(input)
	switch t.Kind() {
	case reflect.Ptr:
		return t.Elem(), nil
	case reflect.Struct:
		return t, nil
	default:
		return nil, errors.New(fmt.Sprintf(InvalidTypeError, t.Kind().String()))
	}
}

func (v *factoryValidator) valueOf(input interface{}) (reflect.Value, error) {
	t := reflect.TypeOf(input)
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.ValueOf(input).Elem(), nil
	case reflect.Struct:
		return reflect.ValueOf(input), nil
	default:
		return reflect.Value{}, errors.New(fmt.Sprintf(InvalidTypeError, t.Kind().String()))
	}
}

func (v *factoryValidator) parseTags(tag string) map[string]string {
	m := make(map[string]string)
	p := `([^\W]+)(=?([^=;]+)?)`
	r, _ := regexp.Compile(p)

	if !r.MatchString(tag) {
		return m
	}

	mm := r.FindAllStringSubmatch(tag, -1)
	for _, sm := range mm {
		m[sm[1]] = sm[3]
	}

	return m
}
