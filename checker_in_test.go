package validation

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleInInput1 struct {
	Age int64 `validate:"in="`
}

type sampleInInput2 struct {
	Age bool `validate:"in=(1,5)"`
}

type sampleInInput3 struct {
	Age int64 `validate:"in=1,5"`
}

type sampleInInput4 struct {
	Age float64 `validate:"in=1.2,5.6"`
}

type sampleInInput5 struct {
	Age string `validate:"in=young,old"`
}

type sampleInInput6 struct {
	Age   int64   `validate:"in=5,10"`
	Money float64 `validate:"in=5.2,10.1"`
	Name  string  `validate:"in=marry,john"`
}

type sampleInInput7 struct {
	Age int64 `validate:"in=1,s"`
}

type sampleInInput8 struct {
	Age float64 `validate:"in=1.2,s"`
}

var _ = Describe("InChecker", func() {
	It("should return error code ERR_VALIDATOR_IN_EMPTY_LIST", func() {
		err := New().Validate(sampleInInput1{10})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(EmptyListFoundError, 10)))
	})

	It("should return error code ERR_VALIDATOR_INVALID_ARGUMENT", func() {
		err := New().Validate(sampleInInput2{false})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(InvalidArgumentError))
	})

	It("should return error code ERR_VALIDATOR_NOT_IN_LIST (int)", func() {
		err := New().Validate(sampleInInput3{7})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(ItemNotFoundInListError, 7, "[1 5]")))
	})

	It("should return error code ERR_VALIDATOR_NOT_IN_LIST (float)", func() {
		err := New().Validate(sampleInInput4{7.3})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(ItemNotFoundInListError, 7.3, "[1.2 5.6]")))
	})

	It("should return error code ERR_VALIDATOR_NOT_IN_LIST (string)", func() {
		err := New().Validate(sampleInInput5{"medium"})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(ItemNotFoundInListError, "medium", "[young old]")))
	})

	It("should return nil", func() {
		err := New().Validate(sampleInInput6{5, 10.1, "john"})
		Expect(err).To(BeNil())
	})

	It("should return error code ERR_VALIDATOR_NOT_INT", func() {
		err := New().Validate(sampleInInput7{5})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(NotIntegerValueError))
	})

	It("should return error code ERR_VALIDATOR_NOT_FLOAT", func() {
		err := New().Validate(sampleInInput8{5})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(NotFloatValueError))
	})
})
