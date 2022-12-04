package dtos_test

import (
	"testing"

	"users/dtos"
)

var validUsernames = []string{
	"username",
	"foo_bar",
	"foo.bar",
	"foo-bar",
	"foo_bar-123",
	"_foo",
}

var invalidUsernames = []string{
	"",
	"fo",
	"foo bar",
	"foo@bar",
	"foo/bar",
	"foo\\bar",
	"foo:bar",
	"12345",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
}

var validEmails = []string{
	"abc-d@mail.com",
	"abc.def@mail.com",
	"abc@mail.com",
	"abc_def@mail.com",
	"abc.def@mail.cc",
	"abc.def@mail-archive.com",
	"abc.def@mail.org",
	"abc.def@mail.com",
}

var invalidEmails = []string{
	"",
	"abc@mail.",
	"abc..def@mail.com",
	".abc@mail.com",
	"abc.def@mail..com",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@mail.com",
}

var validFirstNames = []string{
	"John",
	"Johnathan",
	"Johnathan",
	"Johnathan",
}

var invalidFirstNames = []string{
	"",
	"Johnathannnnnnnnnnnnnnnnnnnnnnn",
}

var validLastNames = []string{
	"John",
	"Johnathan",
	"Johnathan",
	"Johnathan",
}

var invalidLastNames = []string{
	"",
	"Johnathannnnnnnnnnnnnnnnnnnnnnn",
}

func TestRequestDtoValidator(t *testing.T) {
	t.Run("Emails", testEmails)
	t.Run("Usernames", testUsernames)
	t.Run("FirstNames", testFirstNames)
	t.Run("LastNames", testLastNames)
}

func testLastNames(t *testing.T) {
	var usersWithValidLastNames = []dtos.UserRequestDto{}
	var usersWithInvalidLastNames = []dtos.UserRequestDto{}

	for _, lastName := range validLastNames {
		usersWithValidLastNames = append(usersWithValidLastNames, dtos.UserRequestDto{
			Username:  "username",
			Email:     "test@test.com",
			FirstName: "John",
			LastName:  lastName,
		})
	}

	for _, lastName := range invalidLastNames {
		usersWithInvalidLastNames = append(usersWithInvalidLastNames, dtos.UserRequestDto{
			Username:  "username",
			Email:     "test@test.com",
			FirstName: "John",
			LastName:  lastName,
		})
	}

	for _, user := range usersWithValidLastNames {
		var validator = dtos.NewUserValidator(user)
		if !validator.IsValid() {
			t.Errorf("Last Name %s should be valid", user.LastName)
		}
	}

	for _, user := range usersWithInvalidLastNames {
		var validator = dtos.NewUserValidator(user)
		if validator.IsValid() {
			t.Errorf("Last Name %s should be invalid", user.LastName)
		}
	}
}

func testFirstNames(t *testing.T) {
	var usersWithValidFirstNames = []dtos.UserRequestDto{}
	var usersWithInvalidFirstNames = []dtos.UserRequestDto{}

	for _, firstName := range validFirstNames {
		usersWithValidFirstNames = append(usersWithValidFirstNames, dtos.UserRequestDto{
			Username:  "username",
			Email:     "test@test.com",
			FirstName: firstName,
			LastName:  "Doe",
		})
	}

	for _, firstName := range invalidFirstNames {
		usersWithInvalidFirstNames = append(usersWithInvalidFirstNames, dtos.UserRequestDto{
			Username:  "username",
			Email:     "test@test.com",
			FirstName: firstName,
			LastName:  "Doe",
		})
	}

	for _, user := range usersWithValidFirstNames {
		var validator = dtos.NewUserValidator(user)
		if !validator.IsValid() {
			t.Errorf("First name %s should be valid", user.FirstName)
		}
	}

	for _, user := range usersWithInvalidFirstNames {
		var validator = dtos.NewUserValidator(user)
		if validator.IsValid() {
			t.Errorf("First name %s should be invalid", user.FirstName)
		}
	}
}

func testEmails(t *testing.T) {
	var usersWithValidEmails = []dtos.UserRequestDto{}
	var usersWithInvalidEmails = []dtos.UserRequestDto{}

	for _, email := range validEmails {
		usersWithValidEmails = append(usersWithValidEmails, dtos.UserRequestDto{
			Username:  "test",
			Email:     email,
			FirstName: "John",
			LastName:  "Doe"},
		)
	}

	for _, email := range invalidEmails {
		usersWithInvalidEmails = append(usersWithInvalidEmails, dtos.UserRequestDto{
			Username:  "test",
			Email:     email,
			FirstName: "John",
			LastName:  "Doe"},
		)
	}

	for _, user := range usersWithValidEmails {
		var validator = dtos.NewUserValidator(user)
		if !validator.IsValid() {
			t.Errorf("Email %s should be valid", user.Email)
		}
	}

	for _, user := range usersWithInvalidEmails {
		var validator = dtos.NewUserValidator(user)
		if validator.IsValid() {
			t.Errorf("Email %s should be invalid", user.Email)
		}
	}
}

func testUsernames(t *testing.T) {
	var usersWithValidUsernames = []dtos.UserRequestDto{}
	var usersWithInvalidUsernames = []dtos.UserRequestDto{}

	for _, username := range validUsernames {
		usersWithValidUsernames = append(usersWithValidUsernames, dtos.UserRequestDto{
			Username:  username,
			Email:     "test@test.com",
			FirstName: "John",
			LastName:  "Doe",
		})
	}

	for _, username := range invalidUsernames {
		usersWithInvalidUsernames = append(usersWithInvalidUsernames, dtos.UserRequestDto{
			Username:  username,
			Email:     "test@test.com",
			FirstName: "John",
			LastName:  "Doe",
		})
	}

	for _, user := range usersWithValidUsernames {
		var validator = dtos.NewUserValidator(user)
		if !validator.IsValid() {
			t.Errorf("Username %s should be valid", user.Username)
		}
	}

	for _, user := range usersWithInvalidUsernames {
		var validator = dtos.NewUserValidator(user)
		if validator.IsValid() {
			t.Errorf("Username %s should be invalid", user.Username)
		}
	}
}
