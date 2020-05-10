package validation

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validator Suite")
}

type xChecker struct{}

func (c *xChecker) Name() string                              { return "x" }
func (c *xChecker) Check(v interface{}, expects string) error { return nil }

type emptyValidatorInput struct{}
type yChecker struct{}

func (c *yChecker) Name() string                              { return "y" }
func (c *yChecker) Check(v interface{}, expects string) error { return nil }

type zChecker struct{}

func (c *zChecker) Name() string { return "z" }
func (c *zChecker) Check(v interface{}, expects string) error {
	if vv, ok := v.(string); ok == true && vv != "" {
		return nil
	}
	return errors.New("msg")
}

type sampleValidatorInput1 struct {
	Y string `validate:"y=1;z"`
	X int    `not_validate:"true"`
}

type sampleValidatorInput2 struct {
	Y string `validate:""`
}

type sampleValidatorInput3 struct {
	Y int `my_tag:"min=3"`
}

type sampleValidatorInput4 struct {
	Y string `validate:"mailer"`
}

type sampleValidatorInput5 struct {
	X int `validate:"z"`
	Y string
}

type sampleValidatorInput6 struct {
	Z string
	I sampleValidatorInput5
}

type sampleValidatorInput7 struct {
	A int `validate:"max=3" error:"A must be lower than 3"`
}

type sampleValidatorInput8 struct {
	A string `validate:"empty=false"`
	B sampleValidatorInput7
}

type sampleValidatorInput9 struct {
	A map[string]*sampleValidatorInput11 `validate:"nil=false"`
}

type sampleValidatorInput10 struct {
	A []*sampleValidatorInput11 `validate:"nil=false"`
}

type sampleValidatorInput11 struct {
	B string `validate:"minLength=2"`
}

var _ = Describe("Validator", func() {
	It("should return an instance of Validator", func() {
		Expect(New()).NotTo(BeNil())
	})

	It("should return an instance of Validator with custom tag", func() {
		Expect(NewWithTag("my_tag")).NotTo(BeNil())
	})
})
var _ = Describe("factoryValidator", func() {
	It("[private] should parse tags", func() {
		tag := "email;min=30;max=100;in=40,50,55;regexp=(\\d+)"
		v := &factoryValidator{}
		m := v.parseTags(tag)

		Expect(m["in"]).To(Equal("40,50,55"))
		Expect(m["regexp"]).To(Equal(`(\d+)`))
		Expect(m["email"]).To(BeEmpty())
		Expect(m["min"]).To(Equal("30"))
		Expect(m["max"]).To(Equal("100"))
	})

	It("should allow to set Checker", func() {
		v := &factoryValidator{checkers: make(map[string]Checker)}
		Expect(len(v.checkers)).To(BeZero())

		v.Extend(&xChecker{})
		Expect(len(v.checkers)).To(Equal(1))
	})

	It("should return error code ERR_VALIDATOR_INVALID_TYPE", func() {
		err := New().Validate("a string")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(fmt.Sprintf(InvalidTypeError, "string")))
	})

	It("should ignore when no validation rules found", func() {
		err := New().Validate(emptyValidatorInput{})
		Expect(err).To(BeNil())
	})

	It("should validate with custom tag", func() {
		err := NewWithTag("my_tag").Validate(sampleValidatorInput3{Y: 3})
		Expect(err).To(BeNil())
	})

	It("should return nil", func() {
		v := New().Extend(&yChecker{}).Extend(&zChecker{})

		err := v.Validate(sampleValidatorInput1{})
		Expect(err).NotTo(BeNil())

		err = v.Validate(&sampleValidatorInput1{Y: "not_empty"})
		Expect(err).To(BeNil())
	})

	It("should return error 'msg'", func() {
		v := New().Extend(&yChecker{}).Extend(&zChecker{})

		err := v.Validate(sampleValidatorInput6{Z: "a", I: sampleValidatorInput5{1, "b"}})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("X: msg"))
	})

	It("should allow to set custom error message", func() {
		v := New()

		err := v.Validate(sampleValidatorInput7{A: 10})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("A must be lower than 3"))
	})

	It("should validate empty tag", func() {
		v := New().Extend(&yChecker{}).Extend(&zChecker{})
		err := v.Validate(sampleValidatorInput2{"hello"})
		Expect(err).To(BeNil())
	})

	It("should return continue when checker is not found", func() {
		v := &factoryValidator{
			tag:      "validate",
			checkers: make(map[string]Checker),
		}
		err := v.Validate(sampleValidatorInput4{"hello"})
		Expect(err).To(BeNil())
	})

	It("should return parse tags in negative case", func() {
		v := &factoryValidator{}
		tags := v.parseTags("min=2,max=3")
		Expect(len(tags)).To(Equal(2))

		tags = v.parseTags("")
		Expect(len(tags)).To(Equal(0))
	})

	It("should return error if valueOf of wrong type", func() {
		v := &factoryValidator{}
		_, err := v.valueOf("a string")
		Expect(err).NotTo(BeNil())
	})

	It("should test multiple levels", func() {
		v := New()
		err := v.Validate(sampleValidatorInput8{
			A: "OK",
			B: sampleValidatorInput7{A: 4},
		})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("A must be lower than 3"))
	})

	It("should validate Map", func() {
		v := New()
		in := new(sampleValidatorInput9)
		err := v.Validate(in)
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("A", NilValueError)))

		in.A = make(map[string]*sampleValidatorInput11)
		in.A["11"] = &sampleValidatorInput11{B: "A"}
		err = v.Validate(in)
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("B", "minimum length is 2")))

		in.A["11"] = &sampleValidatorInput11{B: "AAA"}
		err = v.Validate(in)
		Expect(err).To(BeNil())
	})

	It("should validate Slice", func() {
		v := New()
		in := new(sampleValidatorInput10)
		err := v.Validate(in)
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("A", NilValueError)))

		in.A = make([]*sampleValidatorInput11, 1)
		in.A[0] = &sampleValidatorInput11{B: "A"}
		err = v.Validate(in)
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("B", "minimum length is 2")))

		in.A[0] = &sampleValidatorInput11{B: "AAA"}
		err = v.Validate(in)
		Expect(err).To(BeNil())
	})
})
