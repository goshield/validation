package validation

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleRangeInputTest1 struct {
	Age string `validate:"range=18-60"`
}

type sampleRangeInputTest2 struct {
	Age int64 `validate:"range=18-60a"`
}

type sampleRangeInputTest3 struct {
	Age int64 `validate:"range=18-60"`
}

type sampleRangeInputTest4 struct {
	Age float64 `validate:"range=18-60a"`
}

type sampleRangeInputTest5 struct {
	Age float64 `validate:"range=10.0-10.2"`
}

type sampleRangeInputTest6 struct {
	Age int64 `validate:"range=9-11"`
}

type sampleRangeInputTest7 struct {
	Age float64 `validate:"range=10.0-10.2"`
}

var _ = Describe("RangeChecker", func() {
	It("should return error code ERR_VALIDATOR_NOT_NUMBER", func() {
		err := New().Validate(sampleRangeInputTest1{"10"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", NotNumberValueError)))
	})

	It("should return error code ERR_VALIDATOR_INVALID_FORMAT", func() {
		err := New().Validate(sampleRangeInputTest2{10})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", InvalidRangeFormatError)))
	})

	It("should return error code ERR_VALIDATOR_NOT_IN_RANGE", func() {
		err := New().Validate(sampleRangeInputTest3{10})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(NotInRangeError, 18, 60))))
	})

	It("should return error code ERR_VALIDATOR_INVALID_FORMAT (float)", func() {
		err := New().Validate(sampleRangeInputTest4{10.2})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", InvalidRangeFormatError)))
	})

	It("should return error code ERR_VALIDATOR_NOT_IN_RANGE (float)", func() {
		err := New().Validate(sampleRangeInputTest5{9.9})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(NotInRangeError, 10.0, 10.2))))
	})

	It("should return nil (int)", func() {
		err := New().Validate(sampleRangeInputTest6{10})
		Expect(err).To(BeNil())
	})

	It("should return nil (float)", func() {
		err := New().Validate(sampleRangeInputTest7{10.1})
		Expect(err).To(BeNil())
	})
})
