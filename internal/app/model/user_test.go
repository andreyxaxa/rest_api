package model_test

import (
	"testing"

	"github.com/andreyxaxa/rest_api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

var (
	tooLongPassword  = "grocolymitxnkjleixgacqixkhdzrlcjqrnraclbmovnppuleswutykruthieazrzyesukwsohrjjfxllqajllevadyoedwfakiq1" // 101 chars
	tooShortPassword = "123"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "too short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = tooShortPassword
				return u
			},
			isValid: false,
		},
		{
			name: "too long password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = tooLongPassword
				return u
			},
			isValid: false,
		},
		{
			name: "with empty password, but with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword123"
				return u
			},
			isValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
