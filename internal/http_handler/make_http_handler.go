package http_handler

import (
	"encoding/json"
	"github.com/kaveesh680/ctrl-print/internal/errors"
	"net/http"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHTTPHandler(f apiFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			e := errors.NewGeneralError(
				http.StatusMethodNotAllowed,
				"oops Something went wrong!",
				"Method Not Allowed",
			)
			err := writeJSON(w, e)
			if err != nil {
				return
			}
			return
		}
		if err := f(w, r); err != nil {
			if e, ok := err.(errors.GeneralError); ok {
				err := writeJSON(w, e)
				if err != nil {
					return
				}
				return
			}
			e := errors.NewGeneralError(
				http.StatusInternalServerError,
				"oops Something went wrong!",
				"Internal Server Error",
			)
			err := writeJSON(w, e)
			if err != nil {
				return
			}
		}
	}
}

func writeJSON(w http.ResponseWriter, e errors.GeneralError) error {
	w.WriteHeader(e.Code)
	w.Header().Set("content-Type", "application/json")
	return json.NewEncoder(w).Encode(e)
}
