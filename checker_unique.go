package validation

import (
	"errors"
	"strings"
)

func UniqueChecker(fetcher DatabaseFetcher) Checker {
	return &uniqueChecker{
		fetcher: fetcher,
	}
}

type uniqueChecker struct {
	fetcher DatabaseFetcher
}

func (c *uniqueChecker) Name() string {
	return "unique"
}

func (c *uniqueChecker) Check(v interface{}, expects string) error {
	var table, column string
	a := strings.Split(expects, ",")
	switch len(a) {
	case 2:
		table = a[0]
		column = a[1]
	default:
		return errors.New(InvalidArgumentError)
	}

	conditions := make(map[string]interface{})
	conditions[column] = v
	row, err := c.fetcher.FetchOne(table, conditions)
	if err != nil {
		return errors.New(UnableToFetchResourceError)
	}
	if row != nil {
		return errors.New(NotUniqueValueError)
	}

	return nil
}
