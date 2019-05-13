package validation

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleDatabaseFetcher struct{}

func (f *sampleDatabaseFetcher) FetchOne(table string, conditions map[string]interface{}) (interface{}, error) {
	if table == "sample_1" {
		return true, nil
	}
	if table == "sample_3" {
		return false, nil
	}
	if table == "sample_4" {
		return nil, errors.New("some error")
	}

	return nil, nil
}

func getUniqueValidator() Validator {
	return New().Extend(UniqueChecker(new(sampleDatabaseFetcher)))
}

type sampleUniqueInput1 struct {
	Email string `validate:"unique"`
}

type sampleUniqueInput2 struct {
	Email string `validate:"unique=sample_1,email"`
}

type sampleUniqueInput3 struct {
	Email string `validate:"unique=sample_2,email"`
}

type sampleUniqueInput4 struct {
	Email string `validate:"unique=sample_3,email"`
}

type sampleUniqueInput5 struct {
	Email string `validate:"unique=sample_4,email"`
}

var _ = Describe("UniqueChecker", func() {
	It("should return error code ERR_VALIDATOR_INVALID_ARGUMENT", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput1{"e@mail.com"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Email", InvalidArgumentError)))
	})

	It("should return error code ERR_VALIDATOR_NOT_UNIQUE", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput2{"e@mail.com"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Email", NotUniqueValueError)))
	})

	It("should return nil", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput3{"e@mail.com"})
		Expect(err).To(BeNil())
	})

	It("should return error code ERR_VALIDATOR_NOT_UNIQUE", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput4{"e@mail.com"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Email", NotUniqueValueError)))
	})

	It("should return error code UnableToFetchResourceError", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput5{"e@mail.com"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Email", UnableToFetchResourceError)))
	})
})
