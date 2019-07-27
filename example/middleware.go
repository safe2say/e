package main

import (
	"e"
	"fmt"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request) error

func withError(next handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, e.ErrorMessage(err), e.HTTPCode(err))
		}

	}
	return fn
}
