package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
	Notes   []string    `json:"notes,omitempty"`
}

type HttpError struct {
	Error string `json:"error"`
}

func NewHttpError(err string) HttpError {
	return HttpError{Error: err}
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

func HttpErrorResponse(err string, notes ...string) Response {
	return NewResponse(NewHttpError(err), notes...)
}

func Abort(c *gin.Context, status int, err string, notes ...string) {
	c.AbortWithStatusJSON(status, HttpErrorResponse(err, notes...))
}

func Ok(c *gin.Context, data interface{}, notes ...string) {
	c.JSON(http.StatusOK, NewResponse(data, notes...))
}

func Err(c *gin.Context, status int, err string, notes ...string) {
	c.JSON(status, HttpErrorResponse(err, notes...))
}

func InternalServerError(c *gin.Context, err string, notes ...string) {
	Abort(c, http.StatusInternalServerError, err, notes...)
}

func Unauthorized(c *gin.Context, err string, notes ...string) {
	Abort(c, http.StatusUnauthorized, err, notes...)
}

func BadRequest(c *gin.Context, err string, notes ...string) {
	Abort(c, http.StatusBadRequest, err, notes...)
}

func NotFound(c *gin.Context, err string, notes ...string) {
	Abort(c, http.StatusNotFound, err, notes...)
}

func Forbidden(c *gin.Context, err string, notes ...string) {
	Abort(c, http.StatusForbidden, err, notes...)
}

type invalidPayloadField struct {
	Field string `json:"field,omitempty"`
	Error string `json:"error,omitempty"`
	Param string `json:"param,omitempty"`
}

func NewInvalidPayloadResponse(err error) Response {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return HttpErrorResponse(err.Error())
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
