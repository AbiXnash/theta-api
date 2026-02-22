package auth

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func formatValidationErrors(err error) []ValidationError {
	var errors []ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			var message string

			switch e.Tag() {
			case "required":
				message = "This field is required"
			case "email":
				message = "Invalid email format"
			case "min":
				message = "Minimum length not satisfied"
			case "len":
				message = "Invalid length"
			case "numeric":
				message = "Must contain only numbers"
			case "oneof":
				message = "Invalid value"
			default:
				message = "Invalid value"
			}

			errors = append(errors, ValidationError{
				Field:   field,
				Message: message,
			})
		}
	}

	return errors
}

func validateBusinessRules(req RegisterRequest) map[string]string {
	errors := make(map[string]string)

	if req.IsSastraStudent {
		if !strings.HasSuffix(strings.ToLower(req.Email), "@sastra.ac.in") {
			errors["email"] = "Please enter a valid SASTRA email (@sastra.ac.in)"
		}
		if len(req.RegNo) != 9 || !isNumeric(req.RegNo) {
			errors["regNo"] = "RegNo must be exactly 9 digits"
		}
	}

	if isBTech(req.Department) {
		if req.Year != "IV" {
			errors["year"] = "BTech students must be in IV year"
		}
	} else {
		if req.Year != "I" && req.Year != "II" && req.Year != "III" {
			errors["year"] = "Only I, II, III year allowed for this department"
		}
	}

	return errors
}

func isBTech(dept string) bool {
	normalized := strings.ToLower(strings.ReplaceAll(dept, " ", ""))
	normalized = strings.ReplaceAll(normalized, ".", "")
	return strings.HasPrefix(normalized, "btech")
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
