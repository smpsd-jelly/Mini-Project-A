package response

import "github.com/go-playground/validator/v10"

func FormatValidationErrors(err error) []FieldError {
	var fieldErrors []FieldError

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok {
		return []FieldError{
			{
				Field:   "request",
				Message: "Invalid request body",
			},
		}
	}

	for _, e := range validationErrors {
		switch e.Field() {

		case "OwnerName":
			fieldErrors = append(fieldErrors, FieldError{
				Field:   "owner_name",
				Message: "Owner Name is required.",
			})

		case "CitizenID":
			if e.Tag() == "required" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "citizen_id",
					Message: "Citizen ID is required.",
				})
			}
			if e.Tag() == "len" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "citizen_id",
					Message: "Citizen ID must be 13 digits.",
				})
			}

		case "PhoneNumber":
			fieldErrors = append(fieldErrors, FieldError{
				Field:   "phone_number",
				Message: "Phone Number is required.",
			})

		case "AccountType":
			if e.Tag() == "required" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "account_type",
					Message: "Account Type is required.",
				})
			}
			if e.Tag() == "oneof" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "account_type",
					Message: "Account Type must be SAVING or CURRENT.",
				})
			}

		case "Balance":
			fieldErrors = append(fieldErrors, FieldError{
				Field:   "balance",
				Message: "Balance must be greater than or equal to 0.",
			})

		case "AccountNumber":
			if e.Tag() == "required" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "account_number",
					Message: "Account Number is required.",
				})
			}

			if e.Tag() == "len" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "account_number",
					Message: "Account Number must be 10 digits.",
				})
			}

			if e.Tag() == "numeric" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "account_number",
					Message: "Account Number format is invalid.",
				})
			}
		case "Amount":
			if e.Tag() == "required" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "amount",
					Message: "Amount is required",
				})
			}
			if e.Tag() == "numeric" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "amount",
					Message: "Amount must be a number",
				})
			}
			if e.Tag() == "gt" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "amount",
					Message: "Amount must be greater than 0",
				})
			}

		case "Description":
			if e.Tag() == "max" {
				fieldErrors = append(fieldErrors, FieldError{
					Field:   "description",
					Message: "Description must not exceed 255 characters",
				})
			}
		}

	}
	return fieldErrors

}
