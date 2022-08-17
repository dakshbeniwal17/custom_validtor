# Custom Validator

Custom validator is your golang struct validator for some of the basic validations. Just pass the validation string in the tags and you are good to go.

## Supported validations

Custom validator currently supports following validations:

| Tag | Validation |
| ------ | ------ |
| is-email | Validates if the given string is a valid email |
| is-phone | Validates if the given string is a valid email |
| regex | Pass in your custom regex to validate against it |

## Installation

```sh
go get github.com/dakshbeniwal17/custom_validator
```

## Usage
Add the validation tags to the struct like this:
```
type myStruct struct {
	Email string `my-validator:"is-email"`
	Phone string `my-validator:"is-phone"`
}
```

And Pass the struct to the `ValidateMyStruct` function.
The function returns a map of strings as key-value pair, where key is the name of the struct key for which the validation failed and value is the validation error. If there are no validation errors then the map returned is empty.

---
****is-email****
```
type myStruct struct {
    Email string `my-validator:"is-email"`
}
```
Validates if the given string value is a valid email or not.
Will pass for the following formats:
- `simple@example.com`
- `very.common@example.com`
- `disposable.style.email.with+symbol@example.com`
- `other.email-with-hyphen@example.com`
- `fully-qualified-domain@example.com`
- `user.name+tag+sorting@example.com`
- `x@example.com`
- `example-indeed@strange-example.com`
- `test/test@test.com`
- `admin@mailserver1`
- `example@s.example`
- `" "@example.org`
- `"john..doe"@example.org`
- `mailhost!username@example.org`
- `"very.(),:;<>[]\".VERY.\"very@\\ \"very\".unusual"@strange.example.com`
- `user%example.com@example.org`
- `user-@example.org`

These are not allowed:
- `Abc.example.com`
- `A@b@c@example.com`
- `a"b(c)d,e:f;g<h>i[j\k]l@example.com`
- `just"not"right@example.com`
- `this is"not\allowed@example.com`
- `this\ still\"not\\allowed@example.com`
- `QA[icon]CHOCOLATE[icon]@test.com`

---
****is-phone****
Validates if the given string value is a valid phone number or not.
```
type myStruct struct {
    Phone string `my-validator:"is-phone"`
}
```
Following formats are allowed:
- `18005551234`
- `1 800 555 1234`
- `+1 800 555-1234`
- `+86 800 555 1234`
- `1-800-555-1234`
- `1 (800) 555-1234`
- `(800)555-1234`
- `(800) 555-1234`
- `(800)5551234`
- `800-555-1234`
- `800.555.1234`
- `800 555 1234x5678`
- `8005551234 x5678`
- `1    800    555-1234`
- `1----800----555-1234`

---
****regex****
This allows you to pass your own regex to match against the value. Pass your regex using  `=`
For example:
```
type myStruct struct {
    Message string `my-validator:"regex=yourRegexHere"`
}
```

---
## Examples

**Basic Validations**

```
package main

import (
	"encoding/json"
	"fmt"

	myValidator "github.com/dakshbeniwal17/custom_validator"
)

type myStruct struct {
	Email string `my-validator:"is-email"`
	Phone string `my-validator:"is-phone"`
}

func main() {
	s := myStruct{
		Email: "testemail@example.com",
		Phone: "0123456789",
	}
	var validationErrors map[string]string
	validationErrors = myValidator.ValidateMyStruct(s)
	if len(validationErrors) > 0 {
		fmt.Println("Errors while validating...")
		json, err := json.Marshal(validationErrors)
		if err != nil {
			fmt.Println(validationErrors)
		}
		fmt.Println(string(json))
		return
	}
	fmt.Println("Validations successful...")
	fmt.Println(len(validationErrors))
}
```

Will return an empty map.
Expected output:
```
Validations successful...
0
```

---
**Failed Validations**

```
package main

import (
	"encoding/json"
	"fmt"

	myValidator "github.com/dakshbeniwal17/custom_validator"
)

type myStruct struct {
	Email string `my-validator:"is-email"`
	Phone string `my-validator:"is-phone"`
}

func main() {
	s := myStruct{
		Email: "this is not an email",
		Phone: "this is not a phone number as well",
	}
	var validationErrors map[string]string
	validationErrors = myValidator.ValidateMyStruct(s)
	if len(validationErrors) > 0 {
		fmt.Println("Errors while validating...")
		json, err := json.Marshal(validationErrors)
		if err != nil {
			fmt.Println(validationErrors)
		}
		fmt.Println(string(json))
		return
	}
	fmt.Println("Validations successful...")
	fmt.Println(len(validationErrors))
}
```

Will return error for both the failed validations.
Expected output:
```
Errors while validating...
{"Email":"`this is not an email` is not a valid email","Phone":"`this is not a phone number as well` is not a valid phone number"}
```

---
**Custom Regex Validations**
```
package main

import (
	"encoding/json"
	"fmt"

	myValidator "github.com/dakshbeniwal17/custom_validator"
)

type myStruct struct {
	Message string `my-validator:"regex=^hello.+"`
}

func main() {
	s := myStruct{
		Message: "hello world",
	}
	var validationErrors map[string]string
	validationErrors = myValidator.ValidateMyStruct(s)
	if len(validationErrors) > 0 {
		fmt.Println("Errors while validating...")
		json, err := json.Marshal(validationErrors)
		if err != nil {
			fmt.Println(validationErrors)
		}
		fmt.Println(string(json))
		return
	}
	fmt.Println("Validations successful...")
	fmt.Println(len(validationErrors))
}
```
This will match the value against the provided regex `^hello.+`.
Will output:
```
Validations successful...
0
```

---
**Failed Custom Regex Validation**
```
package main

import (
	"encoding/json"
	"fmt"

	myValidator "github.com/dakshbeniwal17/custom_validator"
)

type myStruct struct {
	Message string `my-validator:"regex=^hello.+"`
}

func main() {
	s := myStruct{
		Message: "not hello world",
	}
	var validationErrors map[string]string
	validationErrors = myValidator.ValidateMyStruct(s)
	if len(validationErrors) > 0 {
		fmt.Println("Errors while validating...")
		json, err := json.Marshal(validationErrors)
		if err != nil {
			fmt.Println(validationErrors)
		}
		fmt.Println(string(json))
		return
	}
	fmt.Println("Validations successful...")
	fmt.Println(len(validationErrors))
}
```
This will match the value against the provided regex `^hello.+` and the validations will fail.
Will output:
```
Errors while validating...
{"Message":"`not hello world` does not match the given regex: ^hello.+"}
```
## License

MIT

**Free Software, Yeah!**
