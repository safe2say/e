package main

import (
	"e"
)

func createUser(username string) error {
	const op = "main.createUser"
	if username == "nii236" {
		return &e.Error{
			Op:      op,
			Code:    e.ECONFLICT,
			Message: "Username is already in use. Please choose a different username.",
			Meta:    map[string]string{"username": username, "filename": "storage.go"},
		}
	}
	return nil
}
