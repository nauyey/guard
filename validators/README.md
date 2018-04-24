validators: Built-in Validators for Guard
=========================================

Table of Contents
=================

* [Validators Support](#validators-support)
    * [Numeric Validators](#numeric-validators)
    * [String Validators](#string-validators)
    * [Not Nil Validator](#not-nil-validator)
* [Usages](#usages)
* [Roadmap](#roadmap)

---------------------------------------

## Validators Support

### Numeric Validators

* IsOdd
* IsEven
* IntGreaterThan
* IntGreaterThanOrEqualTo
* IntEqualTo
* IntLessThan
* IntLessThanOrEqualTo
* IntInRange

### String Validators

* StringNotBlank
* StringInclusion
* StringExclusion
* StringLength

### Not Nil Validator

* NotNil

---------------------------------------

## Usages

Directly use a validator:
```golang
import "github.com/nauyey/guard/validators"

err := (&validators.IsOdd{Value: 4}).Validate()
// err -> not nil
```

Work with Guard:
```golang
import (
	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
)

type User struct {
	Name   string
	Age    int
	Gender string
	Email  string
}

user := &User{
	Name:   "User Name",
	Age:    10,
	Gender: "",
}

// Validate data
err := guard.Validate(
	&validators.StringNotBlank{Value: user.Name},
	&validators.IntGreaterThan{Value: user.Age, Target: 16}, // invalid Age
	&validators.StringInclusion{Value: user.Gender, In: []string{"female", "male", "other"}}, // invalid Gender
)

if errs, ok := err.(guard.Errors); ok {
	// len(errs.ValidationErrors()) -> 2
}
```

More advanced usages, see [guard documentations](../README.md)

--------------------------------------------------------------

## Roadmap

The following features are what we planed to support. If you are interested in any of them, please send us a pull request. The [How to Contribute](../README.md#how-to-contribute) explains how to send a pull request.

- [ ] Numeric int64
- [ ] Numeric int32
- [ ] Numeric float64
- [ ] Numeric float32
- [ ] String patterns validators, like check email address
- [ ] Alpha
- [ ] Other validators