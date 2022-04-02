package validation

import (
	"udonate/exception"
	"udonate/view_model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Validate(request view_model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Family, validation.Required),
		validation.Field(&request.Username, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
