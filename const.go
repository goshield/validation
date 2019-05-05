package validation

const (
	// Validator errors
	InvalidTypeError     = "%s type is not supported"
	InvalidArgumentError = "invalid argument found"
	NotIntegerValueError = "value is not an integer number"
	NotFloatValueError   = "value is not a float number"
	NotNumberValueError  = "value is not a number"
	NotNilValueError     = "value is not nil"
	NotStringValueError  = "value is not a string"

	// Checkers errors
	InvalidEmailFormatError    = "value is not an email address"
	EmptyListFoundError        = "%v is not expected to appear in an empty list"
	ItemNotFoundInListError    = "%v is not appeared in %v"
	MaxValueError              = "maximum value is %v"
	MaxLengthValueError        = "maximum length is %v"
	MinValueError              = "minimum value is %v"
	MinLengthValueError        = "minimum length is %v"
	InvalidRangeFormatError    = "invalid range format"
	NotInRangeError            = "value must be between %v and %v"
	InvalidRegexPatternError   = "%v is not a valid pattern"
	RegexValueError            = "value does not match %v"
	UnableToFetchResourceError = "unable to fetch appropriate resource"
	NotUniqueValueError        = "value already exists. It must be unique"
)
