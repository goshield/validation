package validation

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleMinInputTest1 struct {
	Age int64 `validate:"min=aa"`
}

type sampleMinInputTest2 struct {
	Age string `validate:"min=10"`
}

type sampleMinInputTest3 struct {
	Age int64 `validate:"min=10"`
}

type sampleMinInputTest4 struct {
	Age float64 `validate:"min=10.2"`
}

type sampleMinInputTest5 struct {
	Age int64 `validate:"min=10"`
}

type sampleMinInputTest6 struct {
	Age float64 `validate:"min=10"`
}

type sampleMinInputTest7 struct {
	Age int64 `validate:"min=10.2"`
}

type sampleMinInputTest8 struct {
	Age float64 `validate:"min=10.1"`
}

var _ = Describe("MinChecker", func() {
	It("should return error code ERR_VALIDATOR_NOT_NUMBER", func() {
		err := New().Validate(&sampleMinInputTest1{Age: 10})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", NotNumberValueError)))

		err = New().Validate(&sampleMinInputTest2{Age: "10"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", NotNumberValueError)))
	})

	It("should return error code ERR_VALIDATOR_NOT_MIN (int)", func() {
		err := New().Validate(&sampleMinInputTest3{Age: 9})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(MinValueError, 10))))
	})

	It("should return nil (float)", func() {
		err := New().Validate(&sampleMinInputTest4{Age: 10.3})
		Expect(err).To(BeNil())
	})

	It("should return nil (int)", func() {
		err := New().Validate(&sampleMinInputTest5{Age: 11})
		Expect(err).To(BeNil())
	})

	It("should return error code ERR_VALIDATOR_NOT_INT", func() {
		err := New().Validate(&sampleMinInputTest6{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", NotIntegerValueError)))
	})

	It("should return error code ERR_VALIDATOR_NOT_FLOAT", func() {
		err := New().Validate(&sampleMinInputTest7{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", NotFloatValueError)))
	})

	It("should return error code ERR_VALIDATOR_NOT_MIN (float)", func() {
		err := New().Validate(&sampleMinInputTest8{Age: 9.9})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(MinValueError, 10.1))))
	})
})
