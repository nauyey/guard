Guard
=====

Guard is a simple, elegant and powerful validation solution for golang.
* **Simple**: Only one main concept -- `Validator` and only one main functional API -- `Validate`.
* **Elegant**: **It doesn't use reflections.** Static type checking. Simple, efficient and readable APIs. 
* **Powerful**: Recursive validations. Built-in validators supplied. Easy to add custom validators. None of differences between custom valdators and built-in validators.

See how easily to use Guard:
```golang
import (
	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
)

name := "User Name"
age := 10
gender := ""

// validate data
err := guard.Validate(
	&validators.StringNotBlank{Value: name},
	&validators.IntGreaterThan{Value: age, Target: 16}, // invalid age
	&validators.StringInclusion{Value: gender, In: []string{"female", "male", "other"}}, // invalid gender
)

if errs, ok := err.(guard.Errors); ok {
	// len(errs.ValidationErrors()) -> 2
}
```

See the power of Guard:
```golang
import (
	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

// Validate implements interface guard.Validator
func (user *User) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: user.Name},
		&validators.IntGreaterThan{Value: user.Age, Target: 16},
	)
}

type Book struct {
	Title  string
	Author *User
}

// Validate implements interface guard.Validator
func (book *Book) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: book.Title},
		book.Author, // here goes the power
	)
}

book := &Book{
	Title: "", // invalid
	Author: &User{
		Name: "User Name",
		Age:  10, // invalid
	},
}

// validate data
err = book.Validate()
if errs, ok := err.(guard.Errors); ok {
	// len(errs.ValidationErrors()) -> 2
}
```

Table of Contents
=================

* [Feature Support](#feature-support)
* [Installation](#installation)
* [Usage](#usage)
    * [Directly Use Validators](#directly-use-validators)
    * [Custom Validator](#custom_validator)
    * [Multiple Validations](#multiple-validations)
    * [Validate Struct](#validate-struct)
    * [Validate Associated Structs](#validate-associated-structs)
    * [Recursive Validations](#recursive-validations)
    * [Strict Validator](#strict-validator)
    * [Allow Nil Validator Instance](#allow-nil-validator-instance)
* [Why Another Valiation Package?](#why-another-valiation-package?)
* [How to Contribute](#how-to-contribute)

---------------------------------------

## Feature Support

* Unified Validator Interface
* Built-in Validators
* Custom Validators
* Override Validation messages
* Strict Validations
* Multiple Validations
* Recursive Validations
* Validate Data Model
* Validate Associated Data Models
* Validation Errors
* Allow Nil Validator

---------------------------------------

## Installation

Simple install the package to your `$GOPATH` with the go tool from shell:

```bash
$ go get -u github.com/nauyey/guard
```

---------------------------------------

## Usage

### Directly Use Validators

The built-in validators in the sub package `"github.com/nauyey/guard/validators"` can be used directly:

```golang
import "github.com/nauyey/guard/validators"

// validate whether a number is odd
v := &validators.IsOdd{Value: 6}
err := v.Validate()
// err -> not nil

// validate string length
v2 := &validators.StringLength{Value: "abc", Min: 2, Max: 7}
err = v.Validate()
// err -> nil
```

More built-in validators, see [validators](./validators/README.md)

### Custom Validator

By implementing interface `"github.com/nauyey/guard".Validator`, anything can be a Validator. They are just the same as the validators from package `"github.com/nauyey/guard/validators"`.

Since the receiver of method `Validator() error` handles most of the complexities, so interface `"github.com/nauyey/guard".Validator` can be so clean.

```golang
type MyValidator struct {
	MyField string
	dbConnection *sql.DB // your validation might need access the database
}

// Validate implements interface guard.Validator
func (user *MyValidator) Validate() error {
	// validation implementations
}
```

### Multiple Validations

```golang
import (
	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
)

// multiple validation
err := guard.Validate(
	&validators.IsOdd{Value: 6},
	&validators.StringNotBlank{Value: " "},
	&validators.StringLength{Value: "abc", Min: 2, Max: 7},
)

// get multiple validation errors
if errs, ok := err.(guard.Errors); ok {
	// len(errs.ValidationErrors()) -> 2
}
```

### Validate Struct

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

// Validate implements interface guard.Validator
func (user *User) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: user.Name},
		&validators.IntGreaterThan{Value: user.Age, Target: 16},
	)
}

user := &User{
	Name: "User Name",
	Age:  10,
}

// validate data
err := user.Validate()
// err -> not nil
```

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
	Name: "User Name",
	Age:  10,
}

// validate data
err := guard.Validate(
	&validators.StringNotBlank{Value: user.Name},
	&validators.IntGreaterThan{Value: user.Age, Target: 16},
	&validators.StringInclusion{Value: user.Gender, In: []string{"female", "male", "other"}},
)
// err -> not nil
```

### Validate Associated Structs

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

// Validate implements interface guard.Validator
func (user *User) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: user.Name},
		&validators.IntGreaterThan{Value: user.Age, Target: 16},
	)
}

type Book struct {
	Title  string
	Author *User
}

// Validate implements interface guard.Validator
func (book *Book) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: book.Title},
		guard.Strict(&validators.NotNil{Value: book.Author}),
		book.Author,
	)
}

book := &Book{
	Title: "", // invalid
	Author: &User{
		Name: "User Name",
		Age:  10, // invalid
	},
}

err = book.Validate()
// err -> not nil
```

### Recursive Validations

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

// Validate implements interface guard.Validator
func (user *User) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: user.Name},
		&validators.IntGreaterThan{Value: user.Age, Target: 16},
	)
}

type Book struct {
	Title  string
	Author *User
}

// Validate implements interface guard.Validator
func (book *Book) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: book.Title},
		guard.Strict(&validators.NotNil{Value: book.Author}),
		book.Author,
	)
}

type Bookmark struct {
	Book    *Book
	Page    int
	Comment string
}

// Validate implements interface guard.Validator
func (b *Bookmark) Validate() error {
	return guard.Validate(
		guard.AllowNil(b.Book),
		&validators.IntGreaterThan{Value: b.Page, Target: 0},
		&validators.StringLength{Value: b.Comment, Min: 0, Max: 255},
	)
}

bookmark := &Bookmark{
	Book: &Book{
		Title: "Book Title",
		Author: &User{
			Name: "User Name",
			Age:  10,
		},
	},
	Page:    10,
	Comment: "hello world!",
}

err := bookmark.Validate()
// err -> not nil
```

### Strict Validator

```golang
import (
	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
)

err := guard.Validate(
	&validators.IsOdd{Value: 6}, // invalid
	guard.Strict(
		&validators.StringNotBlank{Value: " "}, // invalid
		&validators.StringNotBlank{Value: "\n\t"}, // invalid
	),
	&validators.StringLength{Value: "abc", Min: 2, Max: 7},
	guard.Strict(&validators.IntEqualTo{Value: 6, Target: 3}), // invalid
	&validators.StringNotBlank{Value: " "}, // invalid
)

if errs, ok := err.(guard.Errors); ok {
	// len(errs.ValidationErrors()) -> 2 but not 5
}
```

### Allow Nil Validator Instance


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

// Validate implements interface guard.Validator
func (user *User) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: user.Name},
		&validators.IntGreaterThan{Value: user.Age, Target: 16},
	)
}

type Book struct {
	Title  string
	Author *User
}

// Validate implements interface guard.Validator
func (book *Book) Validate() error {
	return guard.Validate(
		&validators.StringNotBlank{Value: book.Title},
		guard.AllowNil(book.Author), // nil book.Author is valid
	)
}

book := &Book{
	Title: "Book Titile",
}

err = book.Validate()
// err -> nil
```

---------------------------------------

## Why Another Valiation Package?

### No Reflection vs Reflection

To quote one of the "Go Proverbs" Rob came up with:
> Clear is better than clever.
>
> Reflection is never clear.

Reflection API loses a lot of valuable type annotating abilities, which may leave bugs to be found until runtime.

So Guard doesn't use reflections.

### Functional API vs Stuct Tag API

Some valiation pacages use struct tag API, like
```golang
type User struct {
	Age       uint8      `validate:"gte=0,lte=130"`
	Email     string     `validate:"required,email"`
	Addresses []*Address `validate:"required,dive,required"`
}
```

Tag API has the following disadvantages:
1. Define a new DSL for validation package. Users have to take time to learn this DSL.
2. No syntax checking. Util runtime, these syntax errors might be found out. **What's worse** is that package [validator](https://github.com/go-playground/validator), [govalidator](https://github.com/asaskevich/govalidator) and [beego/validation](https://github.com/astaxie/beego/tree/master/validation) don't have any ways to check syntax errors even in the runtime.
3. No type checking, too.

### Error Interface vs Error Implementation

It seems it's a common sense to define a validation error type for validation packages. The difference between Guard and other packages is that Guard defines a validation error interface but the other ones define concrete validatoin error types.

Defining a validatoin error interface has the following benefits:
* Custom validation implementations don't directly depend on the validation package, like Guard.
* Abilities of the custom validation error aren't limited by a specific validation package. Take the custom validation error of Guard, it only needs to implement the interface `loupe.Error` or interface `loupe.Errors`. Otherwise, it can define anything it really wants.

### Simple vs Complex

Guard is simple. It only has one core concept `Validator` and one main functional API `Validate`. The other validation packages are complicated:

| Package          | Types | Functional APIs |
| ---------------- | ----- | --------------- |
| Guard | `Validator`, `Error`, `Errors` | `Validate`, `Strict`, `AllowNil` |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | `Validatable`, `Rule`, `skipRule`, `RuleFunc`, `FieldRules`, `ErrFieldPointer`, `ErrFieldNotFound`, `Errors`, `InternalError`, `sql.Valuer`| `Validate`, `ValidateStruct`, `Field` |
| [validator](https://github.com/go-playground/validator) | `FilterFunc`, `CustomTypeFunc`, `TagNameFunc`, `Validate`, `TranslationFunc`, `RegisterTranslationsFunc`, `StructLevelFunc`, `StructLevelFuncCtx`, `StructLevel`, `FieldLevel`, `ValidationErrorsTranslations`, `InvalidValidationError`, `ValidationErrors`, `FieldError`| **Too many complicated APIs** |
| [govalidator](https://github.com/asaskevich/govalidator) | `Validator`, `CustomTypeValidator`, `ParamValidator`, `Errors`, `Error`, `UnsupportedTypeError`, `customTypeTagMap` | `ValidateStruct`, `ErrorByField`, `ErrorsByField`, `SetFieldsRequiredByDefault` |
| [beego/validation](https://github.com/astaxie/beego/tree/master/validation) | `Validator`, `ValidFormer`, `Error`, `Result`, `Validation`, | `Clear`, `HasErrors`, `ErrorMap`, `Error`, `AddError`, `SetError`, `Check`, `Valid`, `RecursiveValid` |

### Valiation Packages Comparing

| Package | No Reflection? | No Tag API? | Functinoal API? | Validation Error? |Built-inn Validators? | Recursive Validations? |
| -- | -- | -- | -- | -- | -- | -- |
| Guard | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | ❌ | ✅ | ✅ | ✅ | ✅ | ✅ |
| [validator](https://github.com/go-playground/validator) | ❌ | ❌ | ❌ | ✅ | ✅ | ✅ |
| [govalidator](https://github.com/asaskevich/govalidator) | ❌ | ❌ | ✅ | ✅ | ✅ | ❌ |
| [beego/validation](https://github.com/astaxie/beego/tree/master/validation) | ❌ | ❌ | ✅ | ✅ | ✅ | ✅ |

-----------------------------------------------------------

## How to Contribute

1. Check for open issues or open a fresh issue to start a discussion around a feature idea or a bug.
2. Fork [the repository](http://github.com/nauyey/guard) on GitHub to start making your changes to the **master** branch (or branch off of it).
3. Write a test which shows that the bug was fixed or that the feature works as expected.
4. Send a pull request and bug the maintainer until it gets merged and published. :) Make sure to add yourself to [AUTHORS](AUTHORS.md).