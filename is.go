package validation

import (
	"errors"
	"reflect"
	"strconv"
)

func IsString(v interface{}) (string, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.String {
		return "", errors.New(NotStringValueError)
	}

	return reflect.ValueOf(v).String(), nil
}

func IsInt(v interface{}) (int64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Int64 {
		return 0, errors.New(NotIntegerValueError)
	}

	return reflect.ValueOf(v).Int(), nil
}

func IsStringInt(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.New(NotIntegerValueError)
	}

	return i, nil
}

func IsFloat(v interface{}) (float64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Float64 {
		return 0.0, errors.New(NotFloatValueError)
	}

	return reflect.ValueOf(v).Float(), nil
}

func IsStringFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, errors.New(NotFloatValueError)
	}

	return f, nil
}

func IsNumber(v interface{}) error {
	_, isF := IsFloat(v)
	_, isI := IsInt(v)
	if isF != nil && isI != nil {
		return errors.New(NotNumberValueError)
	}

	return nil
}

func IsNil(v interface{}) error {
	if v != nil {
		return errors.New(NotNilValueError)
	}

	return nil
}

func IsEmpty(v string) bool {
	return v == ""
}
