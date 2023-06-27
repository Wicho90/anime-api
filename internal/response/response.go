package response

import (
	"net/http"
)

type HttpResponse struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
	Err     string `json:"error"`
}

// BadRequestResponse representa una respuesta de solicitud incorrecta (400).
type BadRequestResponse struct {
	HttpResponse
}

func NewBadRequestResponse(message string, err ...error) *BadRequestResponse {
	badRequest := &BadRequestResponse{
		HttpResponse: HttpResponse{
			Message: message,
			Err:     "Bad Request",
			Code:    http.StatusBadRequest,
		},
	}

	if len(err) > 0 && err[0] != nil {
		badRequest.Err = err[0].Error()
	}

	return badRequest
}

// NotFoundResponse representa una respuesta de recurso no encontrado (404).
type NotFoundResponse struct {
	HttpResponse
}

func NewNotFoundResponse(message string, err ...error) *NotFoundResponse {

	notFound := &NotFoundResponse{
		HttpResponse: HttpResponse{
			Message: message,
			Err:     "Not Found",
			Code:    http.StatusNotFound,
		},
	}

	if len(err) > 0 && err[0] != nil {
		notFound.Err = err[0].Error()
	}

	return notFound
}

// InternalServerErrorResponse representa una respuesta de servidor (500).
type InternalServerErrorResponse struct {
	HttpResponse
}

func NewInternalServerErrorResponse(message string, err ...error) *InternalServerErrorResponse {

	internalServerError := &InternalServerErrorResponse{
		HttpResponse: HttpResponse{
			Message: message,
			Err:     "Internal Server Error",
			Code:    http.StatusInternalServerError,
		},
	}

	if len(err) > 0 && err[0] != nil {
		internalServerError.Err = err[0].Error()
	}

	return internalServerError
}
