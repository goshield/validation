# validation
Tools to validate input

[![Build Status](https://api.travis-ci.org/goshield/validation.svg)](https://travis-ci.org/goshield/validation)

### How To Use
```
package main

import (
	"fmt"
	"github.com/goshield/validation"
)

type Person struct {
	Age int64 `validate:"range=18-60"`
	Name string `validate:"minLength=3;maxLength=10"`
	Email string `validate:"email"`
}

func main() {
	v := validation.New()
	fmt.Println(v.Validate(Person{Age: 17}))
	fmt.Println(v.Validate(Person{Age: 18, Name: "ji", Email: "e@mail.com"}))
	fmt.Println(v.Validate(Person{Age: 18, Name: "jimmy joshua", Email: "e@mail.com"}))
	fmt.Println(v.Validate(Person{Age: 18, Name: "jimmy", Email: "###email.com"}))
	fmt.Println(v.Validate(Person{Age: 19, Name: "jimmy", Email: "e@mail.com"}))
}
```

The result looks like
```
value must be between 18 and 60
minimum length is 3
maximum length is 10
value is not an email address
<nil>
```

### Extends Validator

You could extend validator as much as possible via `Checker`

- First, we need to create a struct which implements Checker interface
- Then we add it into our validator

```
validator.Extend(&MyCustomChecker{})
```

### Customize Error Message

Validator ships default error's message for each checker, it is a simple message. 
However, in some cases you might want to customize your own message which becomes 
more meaning to your customers, you could do as following

```
type Person struct {
	Age int `validate:"min=13" error:"Age must be higher than 13"`
}

func main() {
	v := validation.New()
	fmt.Println(v.Validate(Person{Age: 12}))
}
``` 

Hence, result would be

```
Age must be higher than 13
```

### Internal Checkers

- `EmailChecker` helps to validate if a string is an email address, accepts `string`
```
type User struct {
	Email string `validate:"email"`
}
```

- `InChecker` helps to validate if a value is appeared in a list of values, accepts `string`, `int64` and `float64`
```
type User struct {
	Name string `validate:"in=john,marry,bill"`
}
```

- `MaxChecker` helps to validate if a value is lower than a specific maximum value, accepts `int64` and `float64`
```
type User struct {
	Age int64 `validate:"max=60"`
}
```

- `MaxLengthChecker` helps to validate if a string has length is lower than a specific maximum value, accepts `string`
```
type User struct {
	Username string `validate:"maxLength=100"`
}
```

- `MinChecker` helps to validate if a value is greater than a specific minimum value, accepts `int64` and `float64`
```
type User struct {
	Age int64 `validate:"min=13"`
}
```

- `MinLengthChecker` helps to validate if a string has length is greater than a specific maximum value, accepts `string`
```
type User struct {
	Username string `validate:"minLength=3"`
}
```

- `RangeChecker` helps to validate if a value is in a range, accepts `int64`, `float64`
```
type User struct {
	Age int64 `validate:"range=18-60"`
}
```

- `RegExpChecker` helps to validate if a value matches a proposed regular expression string, accepts `string`
```
type User struct {
	Password string `validate:"regexp=[a-z]{1,}[A-Z]{1,}[0-9]{1,}"`
}
```

- `UniqueChecker` [*] helps to validate if a value exists in database's table, accepts `any`
```
// Add to validator before using
// validator.Extend(UniqueChecker(new(YourDatabaseFetcher)))


type User struct {
	Email string `validate:"unique=users,email"`
}
```

[*] These checkers are not enabled by default, user must specifies it in validator when using

### Flow Chart

![Flow Chart](https://www.lucidchart.com/publicSegments/view/456723cb-80f8-4efd-9e59-2cc5efc408b9/image.png)