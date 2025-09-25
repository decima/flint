package common

import "github.com/go-playground/validator/v10"

type Response struct {
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
	Notes   []string    `json:"notes,omitempty"`
}

func NewResponse(data interface{}, notes ...string) Response {
	message := ""
	if len(notes) >= 1 {
		message = notes[0]
		notes = notes[1:]
	}

	return Response{
		Result:  data,
		Message: message,
		Notes:   notes,
	}
}

type ErrorResponse struct {
	Err error `json:"error"`
}

func NewErrorResponse(err error, details ...string) Response {
	return NewResponse(ErrorResponse{err}, details...)
}

type InvalidPayloadResponse struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type invalidPayloadField struct {
	Field string `json:"field,omitempty"`
	Error string `json:"error,omitempty"`
	Param string `json:"param,omitempty"`
}

func NewInvalidPayloadResponse(err error) Response {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewResponse(err.Error())
	}
	fields := []invalidPayloadField{}
	for _, verr := range validationErrors {
		fields = append(fields, invalidPayloadField{
			Field: verr.Field(),
			Error: verr.Tag(),
			Param: verr.Param(),
		})
	}

	return NewResponse(fields)
}
