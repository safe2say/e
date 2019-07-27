package main

import (
	"e"
	"errors"
	"net/http"
)

func handlerExample(w http.ResponseWriter, r *http.Request) error {
	const op = "main.handlerExample"
	username := r.URL.Query().Get("username")
	if username == "" {
		return &e.Error{
			Op:   op,
			Err:  errors.New("no username provided"),
			Code: e.EINVALID,
		}
	}
	err := createUser(username)
	if err != nil {
		return &e.Error{
			Op:   op,
			Err:  err,
			Code: e.ECONFLICT,
			Meta: map[string]string{"url": r.URL.String()},
		}
	}

	w.Write([]byte("OK"))

	return nil
}
