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

var _ = Describe("InChecker", func() {
	It("should return error code EmptyListFoundError", func() {
		err := New().Validate(sampleInInput1{10})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(EmptyListFoundError, 10))))
	})

	It("should return error code ItemNotFoundInListError (int)", func() {
		err := New().Validate(sampleInInput3{7})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(ItemNotFoundInListError, 7, "[1 5]"))))
	})

	It("should return error code ItemNotFoundInListError (float)", func() {
		err := New().Validate(sampleInInput4{7.3})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(ItemNotFoundInListError, 7.3, "[1.2 5.6]"))))
	})

	It("should return error code ItemNotFoundInListError (string)", func() {
		err := New().Validate(sampleInInput5{"medium"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Age", fmt.Sprintf(ItemNotFoundInListError, "medium", "[young old]"))))
	})

	It("should return nil", func() {
		err := New().Validate(sampleInInput6{5, 10.1, "john"})
		Expect(err).To(BeNil())
	})
})
