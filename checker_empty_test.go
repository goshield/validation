package validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleEmptyInputTest1 struct {
	Name string `validate:"empty=false"`
}

var _ = Describe("EmptyChecker", func() {
	It("should return error code EmptyValueError", func() {
		err := New().Validate(sampleEmptyInputTest1{})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(EmptyValueError))
	})

	It("should return nil", func() {
		err := New().Validate(sampleEmptyInputTest1{
			Name: "John",
		})
		Expect(err).To(BeNil())
	})
})
