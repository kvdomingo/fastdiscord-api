package validation

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Field    string
	Expected string
	Actual   interface{}
}

type XValidator struct {
	validator *validator.Validate
}

type GlobalErrorHandlerResp struct {
	Message string `json:"message"`
}

var (
	validate   = validator.New(validator.WithRequiredStructEnabled())
	xValidator = &XValidator{validator: validate}
)

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := make([]ErrorResponse, 0)

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var el ErrorResponse
			el.Field = err.Field()
			el.Expected = err.Tag() + " " + err.Param()
			el.Actual = err.Value()
			validationErrors = append(validationErrors, el)
		}
	}

	return validationErrors
}

func GetValidatorInstance() *XValidator {
	return xValidator
}
