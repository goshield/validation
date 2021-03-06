package validation

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

// Validator an interface of validator
type Validator interface {
	// Validate checks input value for error
	Validate(v interface{}) error

	// Extend registers a checker
	Extend(checker Checker) Validator
}

// New returns a default Validator
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
		Extend(EmptyChecker()).
		Extend(LengthChecker()).
		Extend(NilChecker())
}

func makeError(name string, err string) error {
	return fmt.Errorf("%s: %s", name, err)
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
		sf := t.Field(i)
		vf := val.Field(i)

		switch vf.Kind() {
		case reflect.Map:
			if vf.IsNil() {
				return v.validate(sf, vf)
			}
			iter := vf.MapRange()
			for iter.Next() {
				in := iter.Value().Interface()
				_, err := v.validateType(in)
				if err != nil {
					// value is not supported
					continue
				}
				err = v.Validate(in)
				if err != nil {
					return err
				}
			}
		case reflect.Slice:
			if vf.IsNil() {
				return v.validate(sf, vf)
			}
			for i := 0; i < vf.Len(); i++ {
				in := vf.Index(i).Interface()
				_, err := v.validateType(in)
				if err != nil {
					// value is not supported
					continue
				}
				err = v.Validate(in)
				if err != nil {
					return err
				}
			}
		case reflect.Struct:
			if err := v.Validate(vf.Interface()); err != nil {
				return err
			}
		case reflect.Ptr:
			if vf.IsNil() {
				return v.validate(sf, vf)
			}
			if err := v.Validate(vf.Interface()); err != nil {
				return err
			}
		default:
			err := v.validate(sf, vf)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (v *factoryValidator) validate(sf reflect.StructField, vf reflect.Value) error {
	if _, ok := sf.Tag.Lookup(v.tag); ok == false {
		return nil
	}

	tag := sf.Tag.Get(v.tag)
	if IsEmpty(tag) {
		return nil
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
			return makeError(sf.Name, err.Error())
		}
	}

	return nil
}

func (v *factoryValidator) validateMap() error {
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
		return nil, fmt.Errorf(InvalidTypeError, t.Kind().String())
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
		return reflect.Value{}, fmt.Errorf(InvalidTypeError, t.Kind().String())
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
