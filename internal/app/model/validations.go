package model

import validation "github.com/go-ozzo/ozzo-validation"

// если у юзера пустой encrypted_password - валидируем password как Required
func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
