package validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleNilInputTest1 struct {
	Name *sampleNilInputTest2 `validate:"nil=false"`
}

type sampleNilInputTest2 struct {
	FirstName string
	LastName  string
}

var _ = Describe("NilChecker", func() {
	It("should return error code NilValueError", func() {
		in := sampleNilInputTest1{}
		in.Name = nil
		err := New().Validate(sampleNilInputTest1{})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Name", NilValueError)))
	})

	It("should return nil", func() {
		err := New().Validate(sampleNilInputTest1{
			Name: &sampleNilInputTest2{
				FirstName: "John",
				LastName:  "Doe",
			},
		})
		Expect(err).To(BeNil())
	})
})
