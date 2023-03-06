package main

import (
	"errors"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, err error, message string, status int) {
	app.logError(r, err)
	http.Error(w, message, status)
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, err, "internal server error", http.StatusInternalServerError)
}

func (app *application) userUnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, errors.New("user unauthorized"), "user unauthorized error", http.StatusUnauthorized)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, errors.New("resource not found"), "resource not found", http.StatusNotFound)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, errors.New("Bad Request"), "bad request", http.StatusBadRequest)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, errors.New("method not allowed"), "method not allowed", http.StatusNotFound)
}
