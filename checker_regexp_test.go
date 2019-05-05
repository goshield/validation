package validation

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleRegexpInputTest1 struct {
	Age int64 `validate:"regexp=**"`
}

type sampleRegexpInputTest2 struct {
	Name string `validate:"regexp=^[a-z0-9]+$"`
}

type sampleRegexpInputTest3 struct {
	Age int64 `validate:"regexp=^\\d+$"`
}

var _ = Describe("RegExpChecker", func() {
	It("should return a checker", func() {
		Expect(RegExpChecker()).NotTo(BeNil())
	})

	It("should return error code ERR_VALIDATOR_REGEXP_WRONG_PATTERN", func() {
		err := New().Validate(sampleRegexpInputTest1{10})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(InvalidRegexPatternError, "**")))
	})

	It("should return error code ERR_VALIDATOR_REGEXP_NOT_MATCH", func() {
		err := New().Validate(sampleRegexpInputTest2{"j@hn"})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(RegexValueError, "^[a-z0-9]+$")))
	})

	It("should return error code ERR_VALIDATOR_NOT_STRING", func() {
		err := New().Validate(sampleRegexpInputTest3{10})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(NotStringValueError))
	})

	It("should return nil", func() {
		err := New().Validate(sampleRegexpInputTest2{"john"})
		Expect(err).To(BeNil())
	})
})
