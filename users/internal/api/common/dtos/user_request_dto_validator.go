package dtos

import (
	"net/mail"
	"regexp"
)

var (
	containsOnlyDigitsUnderscoreDashesDotsAlpha = regexp.MustCompile(`^[0-9a-zA-Z_\.\-]*$`)
	beginsWithUnderscoreOrAlpha                 = regexp.MustCompile(`^[a-zA-Z_].*$`)
)

// UserRequestDtoValidator is a validator for UserRequestDto.
type UserRequestDtoValidator struct {
	UserRequestDto UserRequestDto
	Errors         map[string][]string
}

// NewUserValidator creates a new UserRequestDtoValidator.
func NewUserValidator(u UserRequestDto) *UserRequestDtoValidator {
	return &UserRequestDtoValidator{
		UserRequestDto: u,
		Errors:         make(map[string][]string),
	}
}

// IsValid returns true if the user request dto is valid.
func (uv *UserRequestDtoValidator) IsValid() bool {
	uv.validateUsername()
	uv.validateFirstName()
	uv.validateLastName()
	uv.validateEmail()
	return len(uv.Errors) == 0
}

func (uv *UserRequestDtoValidator) validateUsername() {
	if uv.UserRequestDto.Username == "" {
		uv.Errors["Username"] = append(uv.Errors["Username"], "Username is required")
		return
	}

	if !containsOnlyDigitsUnderscoreDashesDotsAlpha.MatchString(uv.UserRequestDto.Username) {
		uv.Errors["Username"] = append(uv.Errors["Username"], "Username can only contain letters, numbers, dashes, underscores, and dots")
	}

	if !beginsWithUnderscoreOrAlpha.MatchString(uv.UserRequestDto.Username) {
		uv.Errors["Username"] = append(uv.Errors["Username"], "Username must begin with a letter or underscore")
	}

	if len(uv.UserRequestDto.Username) < 3 {
		uv.Errors["Username"] = append(uv.Errors["Username"], "Username must be at least 3 characters long")
	}

	if len(uv.UserRequestDto.Username) > 30 {
		uv.Errors["Username"] = append(uv.Errors["Username"], "Username must be less than 30 characters long")
	}
}

func (uv *UserRequestDtoValidator) validateFirstName() {
	if uv.UserRequestDto.FirstName == "" {
		uv.Errors["FirstName"] = append(uv.Errors["FirstName"], "First name is required")
		return
	}

	if len(uv.UserRequestDto.FirstName) > 30 {
		uv.Errors["FirstName"] = append(uv.Errors["FirstName"], "First name must be less than 30 characters long")
	}
}

func (uv *UserRequestDtoValidator) validateLastName() {
	if uv.UserRequestDto.LastName == "" {
		uv.Errors["LastName"] = append(uv.Errors["LastName"], "Last name is required")
		return
	}

	if len(uv.UserRequestDto.LastName) > 30 {
		uv.Errors["LastName"] = append(uv.Errors["LastName"], "Last name must be less than 30 characters long")
	}
}

func (uv *UserRequestDtoValidator) validateEmail() {
	if uv.UserRequestDto.Email == "" {
		uv.Errors["Email"] = append(uv.Errors["Email"], "Email is required")
		return
	}

	if _, err := mail.ParseAddress(uv.UserRequestDto.Email); err != nil {
		uv.Errors["Email"] = append(uv.Errors["Email"], "Email is invalid")
	}

	if len(uv.UserRequestDto.Email) > 50 {
		uv.Errors["Email"] = append(uv.Errors["Email"], "Email must be less than 50 characters long")
	}
}
