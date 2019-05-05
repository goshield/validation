package validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Is", func() {
	It("IsNil should return error code ERR_VALIDATOR_NOT_NIL", func() {
		err := IsNil(5)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(NotNilValueError))
	})

	It("IsNil should return nil", func() {
		err := IsNil(nil)
		Expect(err).To(BeNil())
	})

	It("IsEmpty should return true/false", func() {
		Expect(IsEmpty("")).To(BeTrue())
		Expect(IsEmpty("some string")).To(BeFalse())
	})
})
