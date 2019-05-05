package validation

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleMinLengthInput1 struct {
	Username int64 `validate:"minLength=3"`
}

type sampleMinLengthInput2 struct {
	Username string `validate:"minLength=aa"`
}

type sampleMinLengthInput3 struct {
	Username string `validate:"minLength=3"`
}

type sampleMinLengthInput4 struct {
	Username string `validate:"minLength=3"`
}

var _ = Describe("MinLengthChecker", func() {
	It("should return error code ERR_VALIDATOR_NOT_STRING", func() {
		err := New().Validate(sampleMinLengthInput1{10})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(NotStringValueError))
	})

	It("should return error code ERR_VALIDATOR_NOT_INT", func() {
		err := New().Validate(sampleMinLengthInput2{"aa"})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(NotIntegerValueError))
	})

	It("should return error code ERR_VALIDATOR_NOT_MIN_LENGTH", func() {
		err := New().Validate(sampleMinLengthInput3{"aa"})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(MinLengthValueError, 3)))
	})

	It("should return nil", func() {
		err := New().Validate(sampleMinLengthInput3{"aaaa"})
		Expect(err).To(BeNil())
	})
})
