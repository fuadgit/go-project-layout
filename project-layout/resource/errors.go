package resource

import (
	"errors"
	"fmt"
	"net/http"

	log "ibfd.org/app/log4u"
)

// TODO Update CustomError name by app specific error name.
// e.g. App name = Tarantula. In that case CustomError name can be renamed to TarantulaError

// CustomError defines errors that are created by ??? resources. // TODO replace ??? by app name
type CustomError struct {
	Status int
	Msg    string
}

// ServerError maps errors to internal server errors.
func ServerError(w http.ResponseWriter, rec *http.Request) {
	if r := recover(); r != nil {
		sendISError(w, fmt.Sprintf("%v", r))
	}
}

// sendISError sends an StatusInternalServerError to the client.
func sendISError(w http.ResponseWriter, msg string) {
	sendError(w, errors.New(msg))
}

// NewError creates a ??? specific error. // TODO replace ??? by app name
func NewError(status int, msg string) *CustomError {
	return &CustomError{Status: status, Msg: msg}
}

func (e *CustomError) Error() string {
	return e.Msg
}

// sendError sends an Error to the client with the defined status if the error
// is a CustomError or else with a status of StatusInternalServerError. // TODO replace CustomError by your defined error name
// If the status is StatusInternalServerError or greater then the error will be logged.
func sendError(w http.ResponseWriter, err error) {
	serr := toCustomError(err)
	if serr.Status >= http.StatusInternalServerError {
		log.Errorln(serr.Error())
	}
	http.Error(w, serr.Error(), serr.Status)
}

func toCustomError(err error) *CustomError {
	if terr, ok := err.(*CustomError); ok {
		return terr
	}
	return NewError(http.StatusInternalServerError, err.Error())
}
